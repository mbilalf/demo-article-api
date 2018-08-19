package repository

import "time"

// TagEntity ...
type TagEntity struct {
	ID   int64  `json: "id"`
	Name string `json: "title"`
}

//ArticleEntity ...
type ArticleEntity struct {
	ID        int64     `json: "id"`
	Title     string    `json: "title"`
	Body      time.Time `json:"body"`
	CreatedAt string    `json:"date"`
}
