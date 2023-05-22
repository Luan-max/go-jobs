package handler

import (
	"net/http"

	"github.com/Luan-max/go-jobs/schemas"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete job
// @Description Delete ajob
// @Tags Jobs
// @Accept json
// @Produce json
// @Param id query string true "Job ID"
// @Success 204 {object} DeleteJobResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /job [delete]

func DeleteJobHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "ID is required param")
		return
	}

	job := schemas.Job{}

	if err := db.First(&job, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Job with ID not found")
		return
	}

	if err := db.Delete(&job).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error to delete JOB")
		return
	}

	sendSuccess(ctx, "delete-job", job, http.StatusNoContent)

}
