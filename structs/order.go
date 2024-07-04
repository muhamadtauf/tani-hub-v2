package structs

import "time"

type Order struct {
	Id          int64         `json:"id"`
	Uuid        string        `json:"uuid"`
	Status      string        `json:"status"`
	Total       float64       `json:"total"`
	UserId      int64         `json:"user_id"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	OrderDetail []OrderDetail `json:"order_detail"`
}
