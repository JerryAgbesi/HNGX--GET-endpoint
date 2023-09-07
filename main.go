package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("api/", getJson)

	port := os.Getenv("PORT")
	fmt.Println()
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

}

type data struct {
	SlackName     string `json:"slack_name"`
	CurrentDay    string `json:"current_day"`
	UtcTime       string `json:"utc_time"`
	Track         string `json:"track"`
	GithubFileUrl string `json:"github_file_url"`
	GithubRepoUrl string `json:"github_repo_url"`
	StatusCode    int    `json:"status_code"`
}

func getJson(c *gin.Context) {
	var response data

	utc_time := time.Now()

	response.UtcTime = utc_time.Format(time.RFC3339)
	response.CurrentDay = utc_time.Format("Monday")

	slack_name := c.DefaultQuery("slack_name", "")
	track := c.DefaultQuery("track", "")

	if slack_name == "" || track == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing or invalid query parameters",
		})
	}

	response.SlackName = slack_name
	response.Track = track

	response.GithubFileUrl = "https://github.com/JerryAgbesi/HNGX--GET-endpoint/blob/main/main.go"
	response.GithubRepoUrl = "https://github.com/JerryAgbesi/HNGX--GET-endpoint"

	response.StatusCode = c.Writer.Status()

	c.IndentedJSON(http.StatusAccepted, response)

}
