package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/mbilalf/demo-article-api/service"
)

type Env struct {
	db *sql.DB
}

func main() {
	// Temp impl - todo: Implement db layer

	service.LoadDummyData()
	//service.SetupDatabase()

	var port = 8000

	router := NewRouter()

	var address = fmt.Sprintf(":%d", port)
	fmt.Println("Running server at ", address)

	log.Fatal(http.ListenAndServe(address, router))
}
