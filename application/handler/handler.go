package handler

import (
	"github.com/Luan-max/go-jobs/infra/config"
	"github.com/Luan-max/go-jobs/infra/database"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = database.GetSQLite()
}
