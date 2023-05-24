package cielo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Luan-max/go-jobs/config"

	"github.com/Luan-max/go-jobs/integrations/helpers"

	"github.com/Luan-max/go-jobs/integrations/cielo/dtos"
)

func CreateCardToken(card dtos.CreditCardDto) (dtos.CardAPIResponse, error) {
	logger := config.GetLogger("transaction")

	requestBody, err := json.Marshal(card)

	if err != nil {
		return dtos.CardAPIResponse{}, errors.New("Error marshaling JSON request body: " + err.Error())
	}

	BASE_URL := "https://apisandbox.cieloecommerce.cielo.com.br/"

	headers := map[string]string{
		"MerchantKey":  "IBGUQWMMADBYBRJJZXTGRFSREJOBNVBNHBYOHNFT",
		"MerchantId":   "aebe297b-17fa-4966-9d12-b75057bcb8fc",
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

	logger.Infof("CIELO:\n", string(body))

	var response dtos.CardAPIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return dtos.CardAPIResponse{}, errors.New("Error unmarshaling JSON: " + err.Error())
	}

	return response, nil
}
