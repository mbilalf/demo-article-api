package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var BaseUrl = "/api/v1/"

var routes = Routes{
	Route{
		"GetAllArticles",
		"GET",
		"articles",
		GetAllArticles,
	},
	Route{
		"SaveArticles",
		"POST",
		"articles",
		SaveArticles,
	},
	Route{
		"GetArticle",
		"GET",
		"articles/{id}",
		GetArticle,
	},

	Route{
		"SearchTags",
		"GET",
		"tags/{tagName}/{date}",
		SearchByTag,
	},
}
