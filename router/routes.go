package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/jobs", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Deu bom mano",
			})
		})
		v1.POST("/jobs", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Deu bom mano",
			})
		})
		v1.PUT("/jobs", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Deu bom mano",
			})
		})
		v1.DELETE("/jobs", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Deu bom mano",
			})
		})
	}
}
