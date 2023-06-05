package cielo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Luan-max/go-jobs/config"

	"github.com/Luan-max/go-jobs/integrations/helpers"

	"github.com/Luan-max/go-jobs/integrations/cielo/dtos"
)

func CreateCardToken(card dtos.CreditCardDto) (dtos.CardAPIResponse, error) {
	logger := config.GetLogger("transaction: Create Card Token")

	requestBody, err := json.Marshal(card)

	if err != nil {
		return dtos.CardAPIResponse{}, errors.New("Error marshaling JSON request body: " + err.Error())
	}

	BASE_URL := os.Getenv("CIELO_URL")

	headers := map[string]string{
		"MerchantKey":  os.Getenv("MERCHANT_KEY"),
		"MerchantId":   os.Getenv("MERCHANT_ID"),
		"Content-Type": "application/json",
	}

	req, err := helpers.JSONRequest(helpers.POST, BASE_URL+"1/card/", requestBody, headers)

	if err != nil {
		return dtos.CardAPIResponse{}, errors.New("Error creating JSON request: " + err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dtos.CardAPIResponse{}, errors.New("Error making JSON request: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtos.CardAPIResponse{}, errors.New("Error reading response body: " + err.Error())
	}

	logger.Infof("CIELO:\n%s", string(body))

	var response dtos.CardAPIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return dtos.CardAPIResponse{}, errors.New("Error unmarshaling JSON: " + err.Error())
	}

	return response, nil
}

func CreatePayment(payment dtos.PaymentRequest) (dtos.TransactionResponse, error) {
	logger := config.GetLogger("transaction: Create Payment")

	requestBody, err := json.Marshal(payment)

	if err != nil {
		return dtos.TransactionResponse{}, errors.New("(createPayment) Error marshaling JSON request body: " + err.Error())
	}

	BASE_URL := os.Getenv("CIELO_URL")

	headers := map[string]string{
		"MerchantKey":  os.Getenv("MERCHANT_KEY"),
		"MerchantId":   os.Getenv("MERCHANT_ID"),
		"Content-Type": "application/json",
	}
	fmt.Print(string(requestBody))
	req, err := helpers.JSONRequest(helpers.POST, BASE_URL+"1/sales/", requestBody, headers)

	if err != nil {
		return dtos.TransactionResponse{}, errors.New("(createPayment) Error creating JSON request: " + err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dtos.TransactionResponse{}, errors.New("(createPayment) Error making JSON request: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtos.TransactionResponse{}, errors.New("(createPayment) Error reading response body: " + err.Error())
	}

	logger.Infof("CIELO CREATE PAYMENT:\n%s", string(body))

	var response dtos.TransactionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return dtos.TransactionResponse{}, errors.New("(createPayment) Error unmarshaling JSON: " + err.Error())
	}

	return response, nil
}
