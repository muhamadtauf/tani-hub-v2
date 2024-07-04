package repository

import (
	"database/sql"
	"tani-hub-v2/structs"
)

func InsertOrder(db *sql.DB, order structs.Order) (err error) {
	sql := "INSERT INTO orders (uuid, status, total, user_id)" +
		" VALUES ($1, $2, $3, $4)"

	errs := db.QueryRow(sql, order.Uuid, order.Status, order.Total, order.UserId)

	for _, orderDetail := range order.OrderDetail {
		sql := "INSERT INTO order_details (price, quantity, total, product_id, order_uuid)" +
			" VALUES ($1, $2, $3, $4, $5)"

		db.QueryRow(sql, orderDetail.Price, orderDetail.Quantity, orderDetail.Total, orderDetail.ProductId, orderDetail.OrderUuid)

		//return errs.Err()
		//if errs != nil {
		//	panic(errs)
		//}
	}

	return errs.Err()
}

func GetAllOrder(db *sql.DB) (err error, results []structs.Order) {
	sql := "SELECT * FROM orders"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order = structs.Order{}

		err = rows.Scan(&order.Id, &order.Uuid, &order.Status, &order.Total, &order.CreatedAt, &order.UpdatedAt, &order.UserId)
		if err != nil {
			panic(err)
		}
		results = append(results, order)
	}
	return
}

func GetOrderByUuid(db *sql.DB, order structs.Order) (err error, results []structs.Order) {
	sql := "SELECT * FROM orders WHERE uuid = $1"
	//sql := "SELECT * FROM orders INNER JOIN order_details ON orders.uuid = order_details.order_uuid WHERE uuid = $1"
	rows, err := db.Query(sql, order.Uuid)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order = structs.Order{}

		err = rows.Scan(&order.Id, &order.Uuid, &order.Status, &order.Total, &order.CreatedAt, &order.UpdatedAt, &order.UserId)
		if err != nil {
			panic(err)
		}
		results = append(results, order)
	}
	return
}
