package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title       string
	Company     string
	Description string
	Remote      bool
	Link        string
	Salary      string
	Benefits    string
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
