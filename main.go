package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Temp impl - todo: Implement db layer
	LoadDummyData()

	var port = 8000

	router := NewRouter()

	var address = fmt.Sprintf(":%d", port)
	fmt.Println("Running server at ", address)

	log.Fatal(http.ListenAndServe(address, router))
}

func LoadDummyData() {
	articles = append(articles, Article{Id: "1", Title: "The Go Getters", Body: "dummy body", Date: time.Now(), Tags: []string{"adventure", "health"}})
	articles = append(articles, Article{Id: "2", Title: "Fast fetchers", Body: "dummy body", Date: time.Now(), Tags: []string{"health", "fun"}})
}
