package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID string  `json:"product_id"`
	Quantity  int     `json: "quantity"`
	Total     float64 `json: "total"`
}
