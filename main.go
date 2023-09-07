package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("api/", getJson)

	// config := cors.DefaultConfig()
    // config.AllowAllOrigins = true
    // config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
    // config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
    // config.ExposeHeaders = []string{"Content-Length"}
    // config.AllowCredentials = true
    // config.MaxAge = 12 * time.Hour

    router.Use(CORS())

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

	

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
	response.CurrentDay = utc_time.Format("Monday") //Specify for format for weekday strings

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

	c.IndentedJSON(http.StatusOK, response)

}
