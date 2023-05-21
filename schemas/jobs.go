package schemas

import (
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
