package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Query string `json:"query"`
}

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
	})
}

func SearchContent(c *gin.Context) {
	var json Message

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println(json.Query)

	c.JSON(200, gin.H{
		"message": json.Query,
	})

}

func main() {
	router := gin.Default()
	router.GET("/health", Health)
	router.POST("/ping", SearchContent)
	router.Run()
}
