package structs

import "time"

type Product struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Code       string    `json:"code"`
	Price      float64   `json:"price"`
	Stock      int64     `json:"stock"`
	IsAtHome   bool      `json:"is_at_home"`
	CategoryId int64     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
