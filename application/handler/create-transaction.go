package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Luan-max/go-jobs/application/integrations/cielo"
	cieloDTO "github.com/Luan-max/go-jobs/application/integrations/cielo/dtos"
	"github.com/Luan-max/go-jobs/application/integrations/helpers"

	"github.com/Luan-max/go-jobs/application/schemas"

	dtos "github.com/Luan-max/go-jobs/application/dtos"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Response    cieloDTO.TransactionResponse
	Transaction schemas.Transaction
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

	card, err := CreateCardtoken(request, ctx)
	if err != nil {
		logger.Errf("error creating card token: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	payment, err := CreatePayment(request, card)
	if err != nil {
		logger.Errf("error creating payment token: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	transaction := schemas.Transaction{
		CardNumber:            request.CardNumber,
		Brand:                 request.CardBrand,
		Month:                 request.ExpirationMonth,
		Year:                  request.ExpirationYear,
		Holder:                request.Holder,
		Status:                payment.Payment.Status,
		ExternalTransactionID: payment.Payment.PaymentId,
		Type:                  payment.Payment.Type,
	}

	if err := db.Create(&transaction).Error; err != nil {
		logger.Errf("error create transaction: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error save transaction in database")
		return
	}

	obj := Response{
		Response:    payment,
		Transaction: transaction,
	}

	sendSuccess(ctx, "transaction processed", obj, http.StatusCreated)
}

func CreateCardtoken(request dtos.CreateTransactionDTO, ctx *gin.Context) (cieloDTO.CardAPIResponse, error) {

	card := cieloDTO.CreditCardDto{
		CustomerName:   request.Holder,
		CardNumber:     request.CardNumber,
		Holder:         request.Holder,
		ExpirationDate: fmt.Sprintf("%s/%s", request.ExpirationMonth, request.ExpirationYear),
		Brand:          request.CardBrand,
	}

	response, err := cielo.CreateCardToken(card)
	if err != nil {
		return cieloDTO.CardAPIResponse{}, err
	}

	return response, nil
}

func CreatePayment(request dtos.CreateTransactionDTO, card cieloDTO.CardAPIResponse) (cieloDTO.TransactionResponse, error) {
	payment := cieloDTO.PaymentRequest{
		Customer: cieloDTO.Customer{
			Name: request.Customer.Name,
		},
		MerchantOrderId: request.OrderID,
		Payment: cieloDTO.PaymentTransactionDTO{
			Type:           request.Type,
			Amount:         request.Amount,
			Installments:   request.Installments,
			SoftDescriptor: "123456789ABCD",
			CreditCard: cieloDTO.CreditCardInfo{
				CardToken:    card.CardToken,
				SecurityCode: request.CVV,
				Brand:        request.CardBrand,
			},
		},
	}

	response, err := cielo.CreatePayment(payment)
	if err != nil {
		return cieloDTO.TransactionResponse{}, err
	}

	return response, nil
}
