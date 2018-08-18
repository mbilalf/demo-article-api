package model

import "time"

type Article struct {
	Id    string    `json: "id"`
	Title string    `json: "title"`
	Date  time.Time `json:"date"`
	Body  string    `json:"body"`
	Tags  []string  `json:"tags"`
}
