package repository

import (
	"database/sql"
	"tani-hub-v2/structs"
)

func GetAllArticle(db *sql.DB) (err error, results []structs.Article) {
	sql := "SELECT id, title, sub_title, is_at_home, created_at, updated_at FROM articles"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var article = structs.Article{}

		err = rows.Scan(&article.Id, &article.Title, &article.SubTitle, &article.IsAtHome, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			panic(err)
		}
		results = append(results, article)
	}
	return
}

func GetArticleById(db *sql.DB, article structs.Article) (err error, results []structs.Article) {
	sql := "SELECT * FROM articles WHERE id = $1"

	rows, err := db.Query(sql, article.Id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var article = structs.Article{}

		err = rows.Scan(&article.Id, &article.Title, &article.SubTitle, &article.Content, &article.IsAtHome, &article.CreatedAt, &article.UpdatedAt)
		if err != nil {
			panic(err)
		}
		results = append(results, article)
	}
	return
}

func InsertArticle(db *sql.DB, articles structs.Article) (err error) {
	sql := "INSERT INTO articles (title, sub_title, content, is_at_home)" +
		" VALUES ($1, $2, $3, $4)"

	errs := db.QueryRow(sql, articles.Title, articles.SubTitle, articles.Content, articles.IsAtHome)

	return errs.Err()
}

func UpdateArticle(db *sql.DB, articles structs.Article) (err error) {
	sql := "UPDATE articles SET title = $1, sub_title = $2, content = $3, is_at_home = $4, updated_at = $5 WHERE id = $6"

	errs := db.QueryRow(sql, articles.Title, articles.SubTitle, articles.Content, articles.IsAtHome, articles.UpdatedAt, articles.Id)

	return errs.Err()
}

func DeleteArticle(db *sql.DB, article structs.Article) (err error) {
	sql := "DELETE FROM articles WHERE id = $1"

	errs := db.QueryRow(sql, article.Id)

	return errs.Err()
}
