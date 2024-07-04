package repository

import (
	"database/sql"
	"tani-hub-v2/structs"
)

func GetAllCategory(db *sql.DB) (err error, results []structs.Category) {
	sql := "SELECT * FROM categories"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}
		results = append(results, category)
	}
	return
}

func GetCategoryById(db *sql.DB, category structs.Category) (err error, results []structs.Category) {
	sql := "SELECT * FROM categories WHERE id = $1"

	rows, err := db.Query(sql, category.Id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}
		results = append(results, category)
	}
	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO categories (name)" +
		" VALUES ($1)"

	errs := db.QueryRow(sql, category.Name)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE categories SET name = $1, updated_at = $2 WHERE id = $3"

	errs := db.QueryRow(sql, category.Name, category.UpdatedAt, category.Id)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM categories WHERE id = $1"

	errs := db.QueryRow(sql, category.Id)

	return errs.Err()
}
