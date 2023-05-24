package dto

import "fmt"

type CreateTransactionDTO struct {
	CardNumber      string `json:"card_number" validate:"required"`
	CVV             string `json:"cvv" validate:"required"`
	ExpirationMonth string `json:"month" validate:"required"`
	ExpirationYear  string `json:"year" validate:"required"`
	CardBrand       string `json:"brand" validate:"required"`
}

func (r *CreateTransactionDTO) Validate() error {
	if r.CardNumber == "" {
		return validatePropsIsRequire("card_number", "string")
	}
	if r.CVV == "" {
		return validatePropsIsRequire("cvv", "string")
	}
	if r.ExpirationMonth == "" {
		return validatePropsIsRequire("expireation month", "string")
	}
	if r.ExpirationYear == "" {
		return validatePropsIsRequire("expiration year", "string")
	}
	if r.CardBrand == "" {
		return validatePropsIsRequire("brand", "string")
	}
	return nil
}

func validatePropsIsRequire(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
