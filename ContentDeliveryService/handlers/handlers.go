package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Repository struct{}

type UserPreference struct {
	Country  string `json:"country"`
	Category string `json:"category"`
}

// define a channel to handle incoming preferences
var preferenceChannel = make(chan UserPreference)

// retrieve User preference and send to the user service
func (m *Repository) GetUserPreference(c *gin.Context) {
	var userRequest UserPreference
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}
	//send the prefernces to the channel asynchronously
	go func(pref UserPreference) {
		preferenceChannel <- pref
	}(userRequest)
	c.JSON(http.StatusOK, gin.H{"message": "Preferences set successfully"})
}

// Process preferences in a separate goroutine
func (m *Repository) StartProcessingPreference() {
	var c *gin.Context
	go func() {
		for {
			select {
			case pref := <-preferenceChannel:
				//process the preference
				err := m.SendUserPreferenceToNewsService(c,pref)
				if err != nil {
					log.Println("Error sending preference to News Service:", err)

				}
			}
		}
	}()
}

// send the preference to the news service
func (m *Repository) SendUserPreferenceToNewsService(c *gin.Context,pref UserPreference) error {
	newServiceUrl := "http//NewsAggregatorServcie/fetchnews"

	jsonData, err := json.MarshalIndent(pref, "", "/t")
	if err != nil {
		log.Println("Unable to marshalIndent preference data", err)
		return err
	}
	//create http client
	client := &http.Client{}

	request, err := http.NewRequest("POST", newServiceUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error in sending prefrence request to newsService", err)
		return err
	}
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

	// Check the HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Println("Unexpected HTTP status code from News Service:", resp.StatusCode)
		return fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	log.Println("Preference sent to News Service:", pref)
	return nil

}
