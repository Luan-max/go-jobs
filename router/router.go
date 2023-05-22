package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	router.Run("0.0.0.0:" + "3000")
}
