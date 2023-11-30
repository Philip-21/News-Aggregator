package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Repo *Repository

type Repository struct{}

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

//Triggers Api fetch the news
func (m *Repository) GetNewsHandler(c *gin.Context) {
	var requestData struct {
		Country  string `json:"country"`
		Category string `json:"category"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}
	news, err := FetchNewsApi(requestData.Country, requestData.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		return
	}

	// Return the aggregated news in the desired format
	c.JSON(http.StatusOK, gin.H{"news": news})
}

// fetch and aggregate news from external APIs
func FetchNewsApi(country, category string) ([]Article, error) {
	apiKey := "8f5aec8663734d3f9f96b2902070889c"

	// Construct the API URL dynamically based on user input
	apiURL := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&category=%s&apiKey=%s", country, category, apiKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Println("unable to make api request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("wrong status code:", err)
		return nil, err
	}

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
