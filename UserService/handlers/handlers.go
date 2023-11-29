package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"user/config"
	"user/database"
	"user/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

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

func (m *Repository) SignUp(c *gin.Context) {
	var RequestPayload database.User
	//Binds JSON request payload which is essential for correctly parsing and handling the incoming data.
	if err := c.ShouldBindJSON(&RequestPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	user, err := m.app.Models.Users.Insert(RequestPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up"})
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
	token, err := middleware.GenerateToken(user.Email, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("token generated")
	_, err = json.Marshal(token)
	if err != nil {
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

func (m *Repository) SetPreference(c *gin.Context) {
	// Extract user ID from the request context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in the request context"})
		return
	}
	var preferences UserPreference
	if err := c.ShouldBindJSON(&preferences); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}
	//store preference in a session
	session := sessions.Default(c)
	session.Set(userID.(string), preferences)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Preferences set successfully"})
	//send the preference to the content delivery service
}
