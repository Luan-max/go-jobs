package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	cieloDTO "github.com/Luan-max/go-jobs/application/integrations/cielo/dtos"
	"github.com/Luan-max/go-jobs/application/integrations/helpers"

	dtos "github.com/Luan-max/go-jobs/application/dtos"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Response    cieloDTO.TransactionResponse
	Transaction any
}

func CreateTransactionHandler(ctx *gin.Context) {
	request := dtos.CreateTransactionDTO{}

	encryptedBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logger.Errf("error reading request body: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "error reading request body")
		return
	}

	decryptedBody, err := helpers.DecryptBody(encryptedBody)
	if err != nil {
		logger.Errf("error decrypting request body: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error decrypting request body")
		return
	}

	if err := json.Unmarshal(decryptedBody, &request); err != nil {
		logger.Errf("error unmarshaling request body: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "error unmarshaling request body")
		return
	}

	ctx.ShouldBindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("error validate request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	card, err := transactionUseCase.CreateCardtoken(&request)
	if err != nil {
		logger.Errf("error creating card token: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	payment, err := transactionUseCase.CreatePayment(request, card)
	if err != nil {
		logger.Errf("error creating payment token: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	createTransactionUseCase, err := transactionUseCase.CreateTransactionUseCase(&request, &payment)
	if err != nil {
		logger.Errf("error creating transaction: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating transaction")
		return
	}

	obj := Response{
		Response:    payment,
		Transaction: createTransactionUseCase,
	}

	sendSuccess(ctx, "transaction processed", obj, http.StatusCreated)
}
