package handler

import (
	"fmt"

	"github.com/Luan-max/go-jobs/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.Status(code)
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}, code int) {
	ctx.Header("Content-type", "application/json")
	ctx.Status(code)
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
		"code":    code,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateTransactionResponse struct {
	Message string                      `json:"message"`
	Data    schemas.TransactionResponse `json:"data"`
	Code    string                      `json:"code"`
}
