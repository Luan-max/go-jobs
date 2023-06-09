package handler

import (
	usecases "github.com/Luan-max/go-jobs/application/usecases/transaction"
	"github.com/Luan-max/go-jobs/infra/config"
	"github.com/Luan-max/go-jobs/infra/database"
	"github.com/Luan-max/go-jobs/infra/database/sqlite/repositories"
	"gorm.io/gorm"
)

var (
	logger             *config.Logger
	db                 *gorm.DB
	transactionRepo    *repositories.SQLiteTransactionRepository
	transactionUseCase *usecases.TransactionUseCaseImpl
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = database.GetSQLite()

	transactionRepo = repositories.NewSQLiteTransactionRepository(db)
	transactionUseCase = usecases.NewTransactionUseCase(transactionRepo)
}
