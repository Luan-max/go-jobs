package repositories

import (
	"github.com/Luan-max/go-jobs/application/schemas"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *schemas.Transaction) error
}

type SQLiteTransactionRepository struct {
	db *gorm.DB
}

func NewSQLiteTransactionRepository(db *gorm.DB) *SQLiteTransactionRepository {
	return &SQLiteTransactionRepository{db: db}
}

func (r *SQLiteTransactionRepository) Create(transaction *schemas.Transaction) error {
	result := r.db.Create(transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
