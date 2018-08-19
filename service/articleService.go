package service

import (
	"fmt"
	"log"
	"time"

	"github.com/mbilalf/demo-article-api/model"
	"github.com/mbilalf/demo-article-api/repository"
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
func SaveArticle(Article *model.Article) (*string, error) {
	db, _ := repository.NewDB()

	tx, err := db.Begin()
	if err != nil {
		log.Println("Unable to start transaction. ", err)
		return nil, err
	}
	_, err = repository.CreateArticleWithTags(db, Article.Title, Article.Body, &Article.Tags)
	if err != nil {
		log.Panic(err)
		err = tx.Rollback()
		if err != nil {
			log.Println("Unable to rollback. ", err)
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Println("Unable to Commit. ", err)
		return nil, err
	}
	var success = ""
	return &success, nil
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

type SearchResult struct {
	Tag         string   `json: "tag"`
	Count       int      `json: "count"`
	Articles    []int    `json: "articles"`
	RelatedTags []string `json: "related_tags"`
}

func SearchArticlesFromDB(Tag string, CreatedAt time.Time) *SearchResult {
	db, _ := repository.NewDB()
	articleSearchData, err := repository.FetchArticlesWithTagAndDate(db, Tag, CreatedAt)
	if err != nil {
		panic(err)
	}

	// Transform result into articleId to tag List map
	artTagListMap := make(map[int][]string)
	for _, searchRec := range articleSearchData {
		fmt.Println("...... ", searchRec.ArticleID, "..", searchRec.TagName)
		artTagListMap[searchRec.ArticleID] = append(artTagListMap[searchRec.ArticleID], searchRec.TagName)
	}

	//Iterate the map to tranform into result instance
	result := SearchResult{Tag: Tag}
	//map to store tags. Map keys will server as a set of tags on all matched articles
	tagMap := make(map[string]bool)
	for artID := range artTagListMap {
		result.Count++
		result.Articles = append(result.Articles, artID)
		for _, t := range artTagListMap[artID] {
			if t != Tag {
				tagMap[t] = true
			}
		}
	}
	for k := range tagMap {
		result.RelatedTags = append(result.RelatedTags, k)
	}
	fmt.Println("...Result... ", result)
	return &result
}

func SetupDatabase() {
	db, _ := repository.NewDB()
	repository.SetupDB(db)
}
