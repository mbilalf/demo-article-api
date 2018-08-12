package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Article struct {
	Id    string    `json: "id"`
	Title string    `json: "title"`
	Date  time.Time `json:"date"`
	Body  string    `json:"body"`
	Tags  []string  `json:"tags"`
}

type SearchResult struct {
	Tag         string   `json: "tag"`
	Count       int      `json: "count"`
	Articles    []string `json: "articles"`
	RelatedTags []string `json: "related_tags"`
}

// Dummy Article cashe - @TODO implement DB layer
var articles []Article

var DateFormat = "2006-01-02"

func SaveArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	// TODO figure out how to set date format while Decoding request
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.Id = strconv.Itoa(rand.Intn(10000000))
	articles = append(articles, article)

	if err := json.NewEncoder(w).Encode(article); err != nil {
		panic(err)
	}
}

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(articles); err != nil {
		panic(err)
	}
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range articles {
		if item.Id == params["id"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				panic(err)
			}
			return
		}
	}

	if err := json.NewEncoder(w).Encode(&Article{}); err != nil {
		panic(err)
	}
}

func HasTag(Target string, AllTags []string) bool {
	for _, tag := range AllTags {
		if tag == Target {
			return true
		}
	}
	return false
}

func MatchDay(targetDate time.Time, date time.Time) bool {
	y, m, d := date.Date()
	y1, m1, d1 := targetDate.Date()
	return y == y1 && m == m1 && d == d1
}

func SearchByTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	tagName := params["tagName"]

	checkDate, _ := time.Parse(DateFormat, params["date"])
	result := SearchResult{Tag: tagName}
	//map to store tags. Map keys will server as a set of tags on all matched articles
	tagMap := make(map[string]bool)

	for _, item := range articles {
		if HasTag(tagName, item.Tags) && MatchDay(checkDate, item.Date) {
			result.Count++
			result.Articles = append(result.Articles, item.Id)
			for _, t := range item.Tags {
				if t != tagName {
					tagMap[t] = true
				}
			}
		}
	}
	//add all Map keys to RelatedTags list
	for k := range tagMap {
		result.RelatedTags = append(result.RelatedTags, k)
	}
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}
