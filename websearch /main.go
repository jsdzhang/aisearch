package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

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
	/**/
	var json Message

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println(json.Query)

	c.JSON(200, gin.H{
		"message": json.Query,
	})
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Hello world")
	}()

}

func main() {
	router := gin.Default()
	router.GET("/health", Health)
	router.POST("/ping", SearchContent)
	router.Run()
}
