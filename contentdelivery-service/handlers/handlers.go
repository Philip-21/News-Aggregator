package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Repo *Repository

type Repository struct{}

type UserPreference struct {
	Country  string `json:"country"`
	Category string `json:"category"`
}
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

// define a channel to handle incoming preferences
var preferenceChannel = make(chan UserPreference)

// retrieve User preference and send to the news service
func (m *Repository) GetUserPreference(c *gin.Context) {
	var userRequest UserPreference
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		log.Println("Invalid JSON format", err)
		return
	}
	err := m.SendUserPreferenceToNewsService(c, userRequest)
	if err != nil {
		log.Println("Unable to send data to news service", err)
		return
	}

	// //send the prefernces to the channel asynchronously
	// go func(pref UserPreference) {
	// 	// Process the preference and send it to the news service
	// 	err := m.SendUserPreferenceToNewsService(c, pref)
	// 	if err != nil {
	// 		log.Println("Unable to send data to news service", err)
	// 	}
	// }(userRequest)
	c.JSON(http.StatusOK, gin.H{"message": "Preferences set successfully"})
	log.Println("News Preferences sent successfully to news service")

}

// send the preference to the news service
func (m *Repository) SendUserPreferenceToNewsService(c *gin.Context, pref UserPreference) error {
	newServiceUrl := "http//newsaggregator-servcie/fetchnews"

	jsonData, err := json.Marshal(pref)
	if err != nil {
		log.Println("Unable to marshal preference data", err)
		return err
	}
	//create http client
	client := &http.Client{}

	request, err := http.NewRequest("POST", newServiceUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error in sending prefrence request to newsService", err)
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	request = request.WithContext(ctx)
	// Use the Gin context's Request object instead of creating a new one
	request = request.WithContext(ctx)
	resp, err := client.Do(request)
	if err != nil {
		log.Println("client request not sent")
		return err
	}
	defer resp.Body.Close()
	log.Println("Preference sent to News Service:", jsonData)

	var respData []Article
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		log.Println("error in decoding", err)
	}
	//m.WriteJSON(c, http.StatusAccepted, respData)
	log.Println("All preference gotten", respData)
	c.JSON(http.StatusAccepted, respData)

	return nil

}

// // NOT IN USE
// Process preferences in a separate goroutine
func (m *Repository) StartProcessingPreference() {
	var c *gin.Context
	go func() {
		for {
			select {
			case pref := <-preferenceChannel:
				//process the preference and send
				err := m.SendUserPreferenceToNewsService(c, pref)
				if err != nil {
					log.Println("Error sending preference to News Service:", err)

				}
			}
		}
	}()
	log.Println("preference requests processed")
}


