package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Luan-max/go-jobs/integrations/cielo"
	cieloDTO "github.com/Luan-max/go-jobs/integrations/cielo/dtos"

	"github.com/Luan-max/go-jobs/schemas"

	dtos "github.com/Luan-max/go-jobs/dtos"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Response    cieloDTO.CardAPIResponse
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

	decryptedBody, err := decryptBody(encryptedBody)
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

	transaction := schemas.Transaction{
		CardNumber: request.CardNumber,
		Brand:      request.CardBrand,
		Month:      request.ExpirationMonth,
		Year:       request.ExpirationYear,
		Holder:     request.Holder,
	}

	card := cieloDTO.CreditCardDto{
		CustomerName:   request.Holder,
		CardNumber:     request.CardNumber,
		Holder:         request.Holder,
		ExpirationDate: fmt.Sprintf("%s/%s", request.ExpirationMonth, request.ExpirationYear),
		Brand:          request.CardBrand,
	}

	response, err := cielo.CreateCardToken(card)
	if err != nil {
		logger.Errf("error creating card token: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Create(&transaction).Error; err != nil {
		logger.Errf("error create transaction: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error save transaction in database")
		return
	}

	obj := Response{
		Response:    response,
		Transaction: transaction,
	}

	sendSuccess(ctx, "transaction processed", obj, http.StatusCreated)
}

func decryptBody(encryptedBody []byte) ([]byte, error) {

	secret := os.Getenv("SECRET")

	key := []byte(secret)

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
