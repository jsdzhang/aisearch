package main

import (
	"encoding/json"
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

type SearchResult struct {
	SearchParameters struct {
		Q      string `json:"q"`
		Type   string `json:"type"`
		Engine string `json:"engine"`
	} `json:"searchParameters"`

	AnswerBox struct {
		Title  string `json:"title"`
		Answer string `json:"answer"`
	} `json:"answerBox"`

	KnowledgeGraph struct {
		Title             string `json:"title"`
		Type              string `json:"type"`
		Website           string `json:"website"`
		ImageURL          string `json:"imageUrl"`
		Description       string `json:"description"`
		DescriptionSource string `json:"descriptionSource"`
		DescriptionLink   string `json:"descriptionLink"`
		Attributes        struct {
			CustomerService string `json:"Customer service"`
			Founders        string `json:"Founders"`
			Founded         string `json:"Founded"`
			Headquarters    string `json:"Headquarters"`
			CEO             string `json:"CEO"`
		} `json:"attributes"`
	} `json:"knowledgeGraph"`

	Organic []struct {
		Title     string `json:"title"`
		Link      string `json:"link"`
		Snippet   string `json:"snippet"`
		Sitelinks []struct {
			Title string `json:"title"`
			Link  string `json:"link"`
		} `json:"sitelinks,omitempty"`
		Date     string `json:"date,omitempty"`
		Position int    `json:"position"`
	} `json:"organic"`

	PeopleAlsoAsk []struct {
		Question string `json:"question"`
		Snippet  string `json:"snippet"`
		Title    string `json:"title"`
		Link     string `json:"link"`
	} `json:"peopleAlsoAsk"`

	RelatedSearches []struct {
		Query string `json:"query"`
	} `json:"relatedSearches"`

	Credits int `json:"credits"`
}

type Result struct {
	Url     string `json:"url"`
	Snippet string `json:"snippet"`
	Content string `json:"content"`
	Summary string `json:"summary"`
}

func GetSearchResult(q string) (*SearchResult, error) {

	url := "https://google.serper.dev/search"
	method := "POST"

	playload := strings.NewReader(`{"q":"` + q + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, playload)

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	req.Header.Add("X-API-KEY", SERPER_API)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	var result SearchResult

	err = json.Unmarshal(body, &result)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return &result, nil
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
