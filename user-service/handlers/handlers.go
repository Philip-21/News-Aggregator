package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"user/config"
	"user/database"
	"user/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type UserPreference struct {
	Country  string `json:"country"`
	Category string `json:"category"`
}
type Repository struct {
	app *config.AppConfig
}

// A constructor function for the Repository struct, it initializes the Repository struct with the app variable,
// making the app accessible within the Repository instance.
// This instance can then be used to initialize and run all your handlers
// This pattern helps in keeping your code modular and makes it easier to manage dependencies
func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{app: app}
}

func (m *Repository) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "welcome to news headline service"})
}

func (m *Repository) GetSignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "create a post request to signup see `https://github.com/Philip-21/News-Aggregator/blob/master/readme.md`"})
}

func (m*Repository)GetSignIn(c*gin.Context){
	c.JSON(http.StatusOK, gin.H{"message":"create a post request to login see `https://github.com/Philip-21/News-Aggregator/blob/master/readme.md`"})
}

func (m *Repository) SignUp(c *gin.Context) {
	var RequestPayload database.User
	//Binds JSON request payload which is essential for correctly parsing and handling the incoming data.
	if err := c.ShouldBindJSON(&RequestPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		log.Println("invalid request payload")
		return
	}
	user, err := m.app.Models.Users.Insert(RequestPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up, User exists"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
	c.JSON(http.StatusOK, gin.H{"message": "User SignedUp successfully"})
	c.Redirect(http.StatusSeeOther, "/user/signin")
	log.Println("signedup successfully")
}

func (m *Repository) Authenticate(c *gin.Context) {
	var UserPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//Binds JSON request payload which is essential for correctly parsing and handling the incoming data.
	if err := c.ShouldBindJSON(&UserPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		log.Println("invalid request payload")
		return
	}
	user, err := m.app.Models.Users.GetEmail(UserPayload.Email)
	if err != nil {
		log.Println("Invalid Email")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid credentials"})
		return
	}
	valid, err := user.PasswordMatches(UserPayload.Password)
	if err != nil || !valid {
		log.Println("Invalid Password")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := middleware.GenerateToken(user.ID, user.Email, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("unable to generate token", err)
		return
	}
	log.Println("token generated")
	_, err = json.Marshal(token)
	if err != nil {
		log.Println("unable to unmarshall data", err)
		return
	}
	log.Println("Token generated:", token)
	c.JSON(http.StatusOK, gin.H{"token": token})
	// Initialize the session and set the userID
	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "User Authenticated successfully"})
	log.Println("Authenticated")
}

// set news prefernce based on country and category
func (m *Repository) SetPreference(c *gin.Context) {
	session := sessions.Default(c)
	userID, exists := c.Get("userID")
	if !exists {
		log.Println("User ID not found in the request context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in the request context"})
		return
	}
	var preferences UserPreference
	if err := c.ShouldBindJSON(&preferences); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}
	session.Set(userID, preferences)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Preferences set successfully"})
	//send the preference to the content delivery service
	err := m.SendPreference(c, "SetNewsPreference", preferences)
	if err != nil {
		return
	}
	log.Println("preference set")
}

// send preference to the content delivery service
func (m *Repository) SendPreference(c *gin.Context, name string, pref UserPreference) error {
	jsonData, err := json.Marshal(pref)
	if err != nil {
		log.Println("Unable to marshal preference data", err)
		return err
	}

	contentServiceUrl := "http://contentdelivery-service/set"

	request, err := http.NewRequest("POST", contentServiceUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in sending Preference request to Content service"})
		log.Println("error in sending Preference request to Content service", err)
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	//cancel the request if the client disconnects
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	request = request.WithContext(ctx)
	// Use the Gin context's Request object instead of creating a new one
	request = request.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println("client request not sent to content service", err)
		return err
	}
	defer resp.Body.Close()
	// // Check the HTTP response status
	// if resp.StatusCode != http.StatusOK {
	// 	log.Println("Unexpected HTTP status code from News Service:", resp.StatusCode)
	// 	return fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	// }
	log.Printf("Preference sent to Content Service: %s", jsonData)
	// rawBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println("error in reading response body")
	// } else {
	// 	log.Printf("Raw Response Body:%s", rawBody)
	// }

	var newsArticle []Article

	err = json.NewDecoder(resp.Body).Decode(&newsArticle)
	if err != nil {
		log.Println("Error in decoding", err)
	}

	//m.WriteJSON(c, http.StatusAccepted, gin.H{"Preferences set successfully": newsArticle})
	c.JSON(http.StatusAccepted, gin.H{"Preferences set successfully": newsArticle})
	return nil
}
