package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var SERPER_API string

type Message struct {
	Query string `json:"query"`
}

func GetSearchResult() {}

// check if the server is healthy
func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
	})
}

/*
The shearch function
*/
func SearchContent(c *gin.Context) {
	/*
		Handle the search here
	*/
	var json Message

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println(SERPER_API)

	c.JSON(200, gin.H{
		"message": json.Query,
	})
	// POST handling
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println(json.Query)
	}()
}

// initial set up , API & DB

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("reading API key fail")
	}
	SERPER_API = os.Getenv("SERPER_API")
	fmt.Println(SERPER_API)

	router := gin.Default()
	router.GET("/health", Health)
	router.POST("/ping", SearchContent)
	router.Run()
}
