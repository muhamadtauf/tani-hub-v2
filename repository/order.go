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
