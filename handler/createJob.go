package handler

import (
	"net/http"

	"github.com/Luan-max/go-jobs/schemas"

	"github.com/gin-gonic/gin"
)

func CreateJobHandler(ctx *gin.Context) {
	request := CreateJobRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("error validate request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	job := schemas.Job{
		Title:       request.Title,
		Company:     request.Company,
		Benefits:    request.Benefits,
		Remote:      *request.Remote,
		Link:        request.Link,
		Description: request.Description,
	}

	if err := db.Create(&job).Error; err != nil {
		logger.Errf("error create job: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error create job in database")
		return
	}

	sendSuccess(ctx, "create-job", job, http.StatusCreated)
}
