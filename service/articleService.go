package service

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/mbilalf/demo-article-api/model"
)

// GetArticles returns all articles
func GetArticles() (*[]model.Article, error) {
	return &articles, nil
}

// GetArticle returns article with given id
func GetArticle(ID string) (*model.Article, error) {
	for _, item := range articles {
		if item.Id == ID {
			return &item, nil
		}
	}
	return nil, nil
}

// SaveArticle creates new article
func SaveArticle(Article *model.Article) (string, error) {
	Article.Id = strconv.Itoa(rand.Intn(10000000))
	articles = append(articles, *Article)
	return Article.Id, nil
}

// SearchArticleByTagAndDate search articles for given tag and date
func SearchArticleByTagAndDate(Tag string, CreatedAt time.Time) (*[]model.Article, error) {

	var matchedArticles []model.Article

	for _, item := range articles {
		if hasTag(Tag, item.Tags) && matchDay(CreatedAt, item.Date) {
			matchedArticles = append(matchedArticles, item)
		}
	}
	return &matchedArticles, nil
}

func hasTag(Target string, AllTags []string) bool {
	for _, tag := range AllTags {
		if tag == Target {
			return true
		}
	}
	return false
}

func matchDay(targetDate time.Time, date time.Time) bool {
	y, m, d := date.Date()
	y1, m1, d1 := targetDate.Date()
	return y == y1 && m == m1 && d == d1
}

var articles []model.Article

func LoadDummyData() {
	articles = append(articles, model.Article{Id: "1", Title: "The Go Getters", Body: "dummy body", Date: time.Now(), Tags: []string{"adventure", "health"}})
	articles = append(articles, model.Article{Id: "2", Title: "Fast fetchers", Body: "dummy body", Date: time.Now(), Tags: []string{"health", "fun"}})
}
