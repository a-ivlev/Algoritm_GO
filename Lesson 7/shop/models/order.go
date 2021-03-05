package models

import "time"

type Order struct {
	ID            int32     `json:"id"`
	CustomerName  string    `json:"customer_name"`
	CustomerPhone string    `json:"customer_phone"`
	CustomerEmail string	`json:"customer_email"`
	ItemIDs       []int32   `json:"item_ids"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}