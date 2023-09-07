package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("api/", getJson)

	router.Run("localhost:9090")

}

type data struct {
	SlackName     string `json:"slack_name"`
	CurrentDay    string `json:"current_day"`
	UtcTime       string `json:"utc_time"`
	Track         string `json:"track"`
	GithubFileUrl string `json:"github_file_url"`
	GithubRepoUrl string `json:"github_repo_url"`
	StatusCode    string `json:"status_code"`
}

// type ctime struct {
// 	Time string `json:"time"`
// }

func getJson(c *gin.Context){
	utc_time := time.Now()
	weekday := utc_time.Format("Monday")


	slack_name := c.DefaultQuery("slack_name","")
	track := c.DefaultQuery("track","")

	if slack_name == "" || track == ""{
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : "Missing or invalid query parameters"
		})
	}
	
	file_url := ""

	

	fmt.Println(slack_name)

	c.IndentedJSON(http.StatusAccepted,weekday)

}