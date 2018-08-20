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
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"date"`
}

// ArticleTagEntity ...
type ArticleTagEntity struct {
	ArticleID int64
	TagID     string
}
