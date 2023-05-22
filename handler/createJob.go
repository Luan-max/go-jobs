package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Luan-max/go-jobs/schemas"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create job
// @Description Create a new job
// @Tags Jobs
// @Accept json
// @Produce json
// @Param request body CreateJobRequest true "Request body"
// @Success 201 {object} CreateJobResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /job [post]

func CreateJobHandler(ctx *gin.Context) {
	request := CreateJobRequest{}

	encryptedBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logger.Errf("error reading request body: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "error reading request body")
		return
	}

	decryptedBody, err := decryptBody(encryptedBody)
	if err != nil {
		logger.Errf("error decrypting request body: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error decrypting request body")
		return
	}

	// BindJSON para ler os dados descriptografados
	if err := json.Unmarshal(decryptedBody, &request); err != nil {
		logger.Errf("error unmarshaling request body: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "error unmarshaling request body")
		return
	}

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

func decryptBody(encryptedBody []byte) ([]byte, error) {

	key := []byte("jobsstudy1234567")

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decodedBody, err := base64.URLEncoding.DecodeString(string(encryptedBody))
	if err != nil {
		return nil, err
	}

	decryptedBody := make([]byte, len(decodedBody)-aes.BlockSize)
	iv := decodedBody[:aes.BlockSize]
	encrypted := decodedBody[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decryptedBody, encrypted)

	return decryptedBody, nil
}
