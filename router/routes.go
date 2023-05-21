package router

import (
	"github.com/Luan-max/go-jobs/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitHandler()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/jobs", handler.CreateJobHandler)
	}
}
