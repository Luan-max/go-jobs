package usecases

import (
	"fmt"

	dto "github.com/Luan-max/go-jobs/application/dtos"
	"github.com/Luan-max/go-jobs/application/integrations/cielo"
	"github.com/Luan-max/go-jobs/application/integrations/cielo/dtos"
	"github.com/Luan-max/go-jobs/application/schemas"
	"github.com/Luan-max/go-jobs/infra/database/sqlite/repositories"
)

type TransactionUseCaseImpl struct {
	transactionRepo repositories.TransactionRepository
}

func NewTransactionUseCase(transactionRepo repositories.TransactionRepository) *TransactionUseCaseImpl {
	return &TransactionUseCaseImpl{
		transactionRepo: transactionRepo,
	}
}

func (uc *TransactionUseCaseImpl) CreateTransactionUseCase(request *dto.CreateTransactionDTO, payment *dtos.TransactionResponse) (*schemas.Transaction, error) {
	transaction := &schemas.Transaction{
		CardNumber:            request.CardNumber,
		Brand:                 request.CardBrand,
		Month:                 request.ExpirationMonth,
		Year:                  request.ExpirationYear,
		Holder:                request.Holder,
		Status:                payment.Payment.Status,
		ExternalTransactionID: payment.Payment.PaymentId,
		Type:                  payment.Payment.Type,
	}

	err := uc.transactionRepo.Create(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (uc *TransactionUseCaseImpl) CreateCardtoken(request *dto.CreateTransactionDTO) (dtos.CardAPIResponse, error) {

	card := dtos.CreditCardDto{
		CustomerName:   request.Holder,
		CardNumber:     request.CardNumber,
		Holder:         request.Holder,
		ExpirationDate: fmt.Sprintf("%s/%s", request.ExpirationMonth, request.ExpirationYear),
		Brand:          request.CardBrand,
	}

	response, err := cielo.CreateCardToken(card)
	if err != nil {
		return dtos.CardAPIResponse{}, err
	}

	return response, nil
}

func (uc *TransactionUseCaseImpl) CreatePayment(request dto.CreateTransactionDTO, card dtos.CardAPIResponse) (dtos.TransactionResponse, error) {
	payment := dtos.PaymentRequest{
		Customer: dtos.Customer{
			Name: request.Customer.Name,
		},
		MerchantOrderId: request.OrderID,
		Payment: dtos.PaymentTransactionDTO{
			Type:           request.Type,
			Amount:         request.Amount,
			Installments:   request.Installments,
			SoftDescriptor: "123456789ABCD",
			CreditCard: dtos.CreditCardInfo{
				CardToken:    card.CardToken,
				SecurityCode: request.CVV,
				Brand:        request.CardBrand,
			},
		},
	}

	response, err := cielo.CreatePayment(payment)
	if err != nil {
		return dtos.TransactionResponse{}, err
	}

	return response, nil
}
