package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Holder     string
	CardNumber string
	Brand      string
	Month      string
	Year       string
}

type TransactionResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt,omitempty"`
	CardNumber string    `json:"card_number"`
	Brand      string    `json:"brand"`
	Month      string    `json:"month"`
	Year       string    `json:"year"`
	Holder     string    `json:"holder"`
}
