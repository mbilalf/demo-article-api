package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

var dateFormat = "2006-01-02"

// ArticleSearchRec ...
type ArticleSearchRec struct {
	ArticleID int
	TagName   string
}

// CreateTag ...
func CreateTag(db *sql.DB, name string) (int64, error) {
	var sql = "INSERT INTO tbl_tag(name) VALUES (?)"
	res, err := db.Exec(sql, name)
	if err != nil {
		log.Println("Error creating Tag. ", err)
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Error geting lastInsertId for tag. ", err)
		return -1, err
	}
	return id, nil
}

//CreateArticle ...
func CreateArticle(db *sql.DB, title string, body string) (int64, error) {
	var sql = `
		INSERT INTO tbl_article(title, body, created_at) 
			VALUES (?, ?, now())
	`
	res, err := db.Exec(sql, title, body)
	if err != nil {
		log.Println("Error creating Article. ", err)
		return -1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Error geting lastInsertId for Article. ", err)
		return -1, err
	}
	return id, nil
}

//CreateArticleTag ...
func CreateArticleTag(db *sql.DB, articleID int64, tagID int64) (bool, error) {
	var sql = `INSERT INTO tbl_article_tag(article_id, tag_id) VALUES (?, ?)`
	_, err := db.Exec(sql, articleID, tagID)
	if err != nil {
		log.Println("Error creating ArticleTag. ", err)
		return false, err
	}
	return true, nil
}

// CreateArticleWithTags ...
func CreateArticleWithTags(db *sql.DB, title string, body string, tags *[]string) (int64, error) {
	//TODO handle duplicate tag
	artID, err := CreateArticle(db, title, body)
	if err != nil {
		return -1, err
	}
	existingTags, err := FetchTagsByName(db, tags)
	if err != nil {
		return -1, err
	}
	var tagIDs []int64
	tagMap := make(map[string]*TagEntity)
	for _, e := range *existingTags {
		tagMap[e.Name] = &e
		tagIDs = append(tagIDs, e.ID)
	}
	for _, t := range *tags {
		if _, ok := tagMap[t]; !ok {
			tagID, err := CreateTag(db, t)
			if err != nil {
				return -1, err
			}
			tagIDs = append(tagIDs, tagID)
		}
	}

	for _, tagID := range tagIDs {
		_, err := CreateArticleTag(db, artID, tagID)
		if err != nil {
			return -1, err
		}
	}
	return artID, nil
}

// FetchTagsByName ...
func FetchTagsByName(db *sql.DB, tags *[]string) (*[]TagEntity, error) {

	tagIs := make([]interface{}, len(*tags))
	for i, t := range *tags {
		tagIs[i] = t
	}

	var sql = "SELECT * FROM tbl_tag WHERE name in (?" + strings.Repeat(",?", len(*tags)-1) + ")"

	rows, err := db.Query(sql, tagIs...)
	if err != nil {
		log.Println("Error fetching tags by name. Error: ", err)
		return nil, err
	}
	var existingTags []TagEntity
	for rows.Next() {
		var t TagEntity
		rows.Scan(&t.ID, &t.Name)
		existingTags = append(existingTags, t)
		fmt.Println("%%%%%", t.ID, "   ", t.Name)
	}
	return &existingTags, nil
}

// FetchAllArticles ...
func FetchAllArticles(db *sql.DB) ([]*ArticleEntity, error) {
	var searchQuery = "SELECT * from tbl_article"
	rows, err := db.Query(searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := make([]*ArticleEntity, 0)
	for rows.Next() {
		art := new(ArticleEntity)
		err := rows.Scan(&art.ID, &art.Title, &art.Body, &art.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, art)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}

// FetchArticlesWithTagAndDate ...
func FetchArticlesWithTagAndDate(db *sql.DB, tagName string, createdAt time.Time) ([]*ArticleSearchRec, error) {
	//
	var queryFormat = `
	SELECT art.id, tbl_tag.name 
	FROM tbl_article art 
	INNER JOIN tbl_article_tag on art.id = tbl_article_tag.article_id
	INNER JOIN tbl_tag on tbl_article_tag.tag_id = tbl_tag.id
	where art.id in (
			SELECT a.id
			FROM tbl_article a
			INNER JOIN tbl_article_tag atag ON a.id = atag.article_id 
			INNER JOIN tbl_tag t ON t.id = atag.tag_id and t.name = '%s'
			WHERE DATE(a.created_at) = '%s')
	ORDER BY art.id;
	`
	dateStr := createdAt.Format(dateFormat)
	var searchQuery = fmt.Sprintf(queryFormat, tagName, dateStr)

	rows, err := db.Query(searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	artSearchData := make([]*ArticleSearchRec, 0)
	for rows.Next() {
		searchRec := new(ArticleSearchRec)
		err := rows.Scan(&searchRec.ArticleID, &searchRec.TagName)
		if err != nil {
			return nil, err
		}
		artSearchData = append(artSearchData, searchRec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return artSearchData, nil
}
