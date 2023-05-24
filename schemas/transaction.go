package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CardNumber string
	Brand      string
	Month      string
	Year       string
}

type JobResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Description string    `json:"description"`
	Remote      bool      `json:"remote"`
	Link        string    `json:"link"`
	Salary      string    `json:"salary"`
	Benefits    string    `json:"benefits"`
}
