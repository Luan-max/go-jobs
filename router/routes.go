package router

import (
	docs "github.com/Luan-max/go-jobs/docs"
	"github.com/Luan-max/go-jobs/handler"
	"github.com/Luan-max/go-jobs/interceptor"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitHandler()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		v1.Use(interceptor.EncryptInterceptor())
		v1.POST("/job", handler.CreateJobHandler)
		v1.DELETE("/job", handler.DeleteJobHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
