package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mbilalf/demo-article-api/service"
)

func main() {
	// Temp impl - todo: Implement db layer
	service.LoadDummyData()

	var port = 8000

	router := NewRouter()

	var address = fmt.Sprintf(":%d", port)
	fmt.Println("Running server at ", address)

	log.Fatal(http.ListenAndServe(address, router))
}
