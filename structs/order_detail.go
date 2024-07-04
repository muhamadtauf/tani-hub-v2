package structs

import "time"

type OrderDetail struct {
	Id        int64     `json:"id"`
	Price     float64   `json:"price"`
	Quantity  int64     `json:"quantity"`
	Total     float64   `json:"total"`
	ProductId int64     `json:"product_id"`
	OrderCode string    `json:"order_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
