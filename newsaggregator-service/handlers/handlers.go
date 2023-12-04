package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type NewsAPIResponse struct {
	Articles []Article `json:"articles"`
}

// func (m *Repository) GetNewsHandler(c *gin.Context) {
// 	///retrive query parameters
// 	country := c.Query("country")
// 	category := c.Query("category")
// 	news, err := FetchNewsApi(country, category)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
// 		return
// 	}
// 	// Return the aggregated news in the desired format
// 	c.JSON(http.StatusOK, gin.H{"news": news})
// }

// Triggers Api fetch the news
func (m *Repository) GetNewsHandler(c *gin.Context) {
	// Read and log the raw request body
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
	} else {
		log.Printf("Raw Request Body: %s", rawBody)
	}
	// Rewind the request body so it can be read again
	c.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	var rawRequest struct {
		Country  string `json:"country"`
		Category string `json:"category"`
	}
	// Extract values from raw JSON string
	if err := json.Unmarshal(rawBody, &rawRequest); err != nil {
		log.Println("Error unmarshalling raw request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error unmarshalling raw request body"})
		return
	}

	// Use extracted values directly or create a new UserPreference struct
	// and assign the values to it.
	requestData := UserPreference{
		Country:  rawRequest.Country,
		Category: rawRequest.Category,
	}

	log.Println("Preferences received from Content Service:", requestData)
	log.Println("Country:", rawRequest.Country)
	log.Println("Category:", rawRequest.Category)

	// Use the extracted values in FetchNewsApi
	news, err := m.FetchNewsApi(requestData.Country, requestData.Category)
	if err != nil {
		log.Println("failed to fetch news from News Api", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		return
	}
	c.JSON(http.StatusAccepted, news)
	log.Println("news generated", news)
}

// fetch and aggregate news from external APIs
func (m *Repository) FetchNewsApi(country, category string) ([]Article, error) {
	//	var c *gin.Context
	apiKey := "8f5aec8663734d3f9f96b2902070889c"

	// Construct the API URL dynamically based on user input
	apiURL := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&category=%s&apiKey=%s", country, category, apiKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Println("unable to make api request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var newsResponse NewsAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&newsResponse); err != nil {
		log.Println("unable to decode jsonData:", err)
		return nil, err
	}
	// Extract relevant information from the API response
	var aggregatedNews []Article
	for _, article := range newsResponse.Articles {
		aggregatedNews = append(aggregatedNews, Article{
			Source: Source{
				ID:   article.Source.ID,
				Name: article.Source.Name,
			},
			Author:      article.Author,
			Title:       article.Title,
			Description: article.Description,
			PublishedAt: article.PublishedAt,
			Content:     article.Content,
		})
	}
	log.Println("Articles gotten successfully")

	return aggregatedNews, nil
}


