package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mbilalf/demo-article-api/model"
	"github.com/mbilalf/demo-article-api/service"
)

type SearchResult struct {
	Tag         string   `json: "tag"`
	Count       int      `json: "count"`
	Articles    []string `json: "articles"`
	RelatedTags []string `json: "related_tags"`
}

var DateFormat = "2006-01-02"

func SaveArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article model.Article
	//TODO figure out how to set date format while Decoding request
	_ = json.NewDecoder(r.Body).Decode(&article)

	service.SaveArticle(&article)
	if err := json.NewEncoder(w).Encode(article); err != nil {
		panic(err)
	}
}

// GetAllArticles ...
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	articles, serviceErr := service.GetArticles()
	if serviceErr != nil {
		log.Println("Failed to get All Articles from service: ", serviceErr)
		if err := json.NewEncoder(w).Encode(serviceErr); err != nil {
			panic(err)
		}
	}
	if err := json.NewEncoder(w).Encode(*articles); err != nil {
		panic(err)
	}
}

// GetArticle ...
func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	article, serviceErr := service.GetArticle(params["id"])
	if serviceErr != nil {
		log.Println("Failed to get Article from service: ", serviceErr)
		if err := json.NewEncoder(w).Encode(serviceErr); err != nil {
			panic(err)
		}
	}
	if err := json.NewEncoder(w).Encode(*article); err != nil {
		panic(err)
	}
}

// SearchByTag ...
func SearchByTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	tagName := params["tagName"]
	checkDate, _ := time.Parse(DateFormat, params["date"])

	result := SearchResult{Tag: tagName}
	//map to store tags. Map keys will server as a set of tags on all matched articles
	tagMap := make(map[string]bool)

	// get search result from service and transform as per requirements
	articles, _ := service.SearchArticleByTagAndDate(tagName, checkDate)
	for _, item := range *articles {
		result.Count++
		result.Articles = append(result.Articles, item.Id)
		for _, t := range item.Tags {
			if t != tagName {
				tagMap[t] = true
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
