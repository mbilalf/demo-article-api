package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mbilalf/demo-article-api/service"
	"github.com/mbilalf/demo-article-api/view"
)

var DateFormat = "2006-01-02"

func SetupDb(w http.ResponseWriter, r *http.Request) {
	service.SetupDatabase()
}

// GetArticle ...
func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusNotImplemented)
	errRes := view.ErrorResponse{Code: http.StatusNotImplemented, Message: fmt.Sprintf("Operation not supported")}
	if err := json.NewEncoder(w).Encode(errRes); err != nil {
		panic(err)
	}
	// TODO add support in repository to enable following code
	/*
		params := mux.Vars(r)
		idStr := params["id"]
		ID, err := strconv.ParseInt(idStr, 10, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errRes := view.ErrorResponse{Code: 400, Message: fmt.Sprintf("Invalid value %s. Expected int", idStr)}
			if err := json.NewEncoder(w).Encode(errRes); err != nil {
				panic(err)
			}
			return
		}
		article, serviceErr := service.GetArticle(ID)
		if serviceErr != nil {
			log.Println("Failed to get Article from service: ", serviceErr)
			if err := json.NewEncoder(w).Encode(serviceErr); err != nil {
				panic(err)
			}
		}

		if err := json.NewEncoder(w).Encode(*article); err != nil {
			panic(err)
		}
	*/
}

// GetAllArticles ... Better implementation would be to paginate instead of fetch all
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

//SaveArticles ...
func SaveArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article view.Article
	//TODO figure out how to set date format while Decoding request
	_ = json.NewDecoder(r.Body).Decode(&article)

	artID, err := service.SaveArticle(&article)
	if err != nil {
		errRes := view.ErrorResponse{Code: http.StatusInternalServerError, Message: fmt.Sprintf("Unable to save article")}
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			panic(err)
		}
		log.Panic("Unable to Save Article. Error: ", err)
	}

	//TODO return saved article record instead of received request object
	article.Id = artID
	if err := json.NewEncoder(w).Encode(article); err != nil {
		panic(err)
	}
}

// SearchByTag ...
func SearchByTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	tagName := params["tagName"]
	checkDate, _ := time.Parse(DateFormat, params["date"])
	result := service.SearchArticlesFromDB(tagName, checkDate)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}
