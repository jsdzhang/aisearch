package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

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

	fmt.Println(playload)

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
	GetSearchResult("apple inc")

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	c.JSON(200, gin.H{
		"message": json.Query,
	})
	// POST handling
	go func() {
		time.Sleep(5 * time.Second)
		//fmt.Println(json.Query)
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
	fmt.Println(SERPER_API)

	router := gin.Default()
	router.GET("/health", Health)
	router.POST("/ping", SearchContent)
	router.Run()
}
