package view

import "time"

type Article struct {
	Id    int64     `json: "id"`
	Title string    `json: "title"`
	Body  string    `json:"body"`
	Date  time.Time `json:"date"`
	Tags  []string  `json:"tags"`
}

type ErrorResponse struct {
	Code    int64  `json: "code"`
	Message string `json: "message"`
}
