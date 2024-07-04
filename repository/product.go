package repository

import (
	"database/sql"
	"tani-hub-v2/structs"
)

func GetAllProduct(db *sql.DB) (err error, results []structs.Product) {
	sql := "SELECT * FROM products"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var product = structs.Product{}

		err = rows.Scan(&product.Id, &product.Name, &product.Code, &product.Price, &product.Stock, &product.IsAtHome, &product.CategoryId, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			panic(err)
		}
		results = append(results, product)
	}
	return
}

func GetProductById(db *sql.DB, product structs.Product) (err error, results []structs.Product) {
	sql := "SELECT * FROM products WHERE id = $1"

	rows, err := db.Query(sql, product.Id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var product = structs.Product{}

		err = rows.Scan(&product.Id, &product.Name, &product.Code, &product.Price, &product.Stock, &product.IsAtHome, &product.CategoryId, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			panic(err)
		}
		results = append(results, product)
	}
	return
}

func InsertProduct(db *sql.DB, product structs.Product) (err error) {
	sql := "INSERT INTO products (name, code , price, stock, is_at_home, category_id)" +
		" VALUES ($1, $2, $3, $4, $5, $6)"

	errs := db.QueryRow(sql, product.Name, product.Code, product.Price, product.Stock, product.IsAtHome, product.CategoryId)

	return errs.Err()
}

func UpdateProduct(db *sql.DB, product structs.Product) (err error) {
	sql := "UPDATE products SET name = $1, code = $2, price = $3, stock = $4, is_at_home = $5, category_id = $6, updated_at = $7 WHERE id = $8"

	errs := db.QueryRow(sql, product.Name, product.Code, product.Price, product.Stock, product.IsAtHome, product.CategoryId, product.UpdatedAt, product.Id)

	return errs.Err()
}

func DeleteProduct(db *sql.DB, product structs.Product) (err error) {
	sql := "DELETE FROM products WHERE id = $1"

	errs := db.QueryRow(sql, product.Id)

	return errs.Err()
}
