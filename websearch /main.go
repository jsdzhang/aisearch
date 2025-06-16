package main

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var SERPER_API string

type Message struct {
	Query string `json:"query"`
}

type Result struct {
	Url     string `json:"url"`
	Snippet string `json:"snippet"`
	Content string `json:"content"`
	Summary string `json:"summary"`
}

func GetSearchResult(q string) (string, error) {

	url := "https://google.serper.dev/search"
	method := "POST"
	playload := strings.NewReader(`{"q":"` + q + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, playload)

	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	req.Header.Add("X-API-KEY", SERPER_API)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		slog.Error(err.Error())
		return "", err
	}

	return string(body), nil
}

func AddSearchResult() {}

/*
The shearch function
*/
func SearchContent(c *gin.Context) {
	/*
		Handle the search here
	*/
	var json Message
	//search query handler

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	res, err := GetSearchResult(json.Query)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	c.JSON(200, gin.H{
		"message":  json.Query,
		"resposne": res,
	})
	// POST handling
	go func() {
		//handle the res here
	}()
}

// initial set up , API & DB

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("reading API key fail")
	}
	SERPER_API = os.Getenv("SERPER_API")

	router := gin.Default()
	router.GET("/health", Health)
	router.POST("/ping", SearchContent)
	router.Run()
}
