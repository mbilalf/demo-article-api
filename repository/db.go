package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {

	db, err := sql.Open("mysql",
		"root:fidodido@tcp(127.0.0.1:3306)/demo_db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func SetupDB(db *sql.DB) {
	const createTableArticle = `
		CREATE TABLE IF NOT EXISTS tbl_article (
			id BIGINT(20) NOT NULL AUTO_INCREMENT, 
			title VARCHAR(255)  NOT NULL, 
			body TEXT, 
			created_at DATETIME NOT NULL,
			PRIMARY KEY (id)
		);
	`
	const createTableTag = `
		CREATE TABLE IF NOT EXISTS tbl_tag (
			id BIGINT(20) NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL UNIQUE,
			PRIMARY KEY (id)
		);
	`
	const createTableArticleTag = `
		CREATE TABLE IF NOT EXISTS tbl_article_tag(
			article_id BIGINT(20) NOT NULL,
			tag_id BIGINT(20) NOT NULL,
			PRIMARY KEY (article_id, tag_id)
		);
	`
	const alterTableArticleTag = `
		ALTER TABLE tbl_article_tag
			ADD FOREIGN KEY (article_id) REFERENCES tbl_article(id),
			ADD FOREIGN KEY (tag_id) REFERENCES tbl_tag(id);
	`
	const insertArticleData = `
		INSERT INTO tbl_article(title, body, created_at) 
		VALUES ("The Go Getters", "dummy body...", now()),
			("Fast Fetchers", "dummy body...", now()),
			("My feet and the trail", "dummy body...", now());
	`
	const insertTagData = `
		INSERT INTO tbl_tag(name) 
			VALUES ("health"), ("adventure"), ("fun");
	`
	const insertArticleTagData = `
		INSERT INTO tbl_article_tag(article_id, tag_id) 
			VALUES (1, 1), (1,2), (2,1), (3,2), (3,3);
	`

	_, err := db.Exec(createTableArticle)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(createTableTag)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(createTableArticleTag)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(alterTableArticleTag)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(insertArticleData)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(insertTagData)
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(insertArticleTagData)
	if err != nil {
		log.Panic(err)
	}

}
