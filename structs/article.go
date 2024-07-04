package structs

import "time"

type Article struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	SubTitle  string    `json:"sub_title"`
	Content   string    `json:"content"`
	IsAtHome  bool      `json:"is_at_home"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
