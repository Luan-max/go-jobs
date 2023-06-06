package dtos

type Customer struct {
	Name string `json:"Name"`
}

type CreditCardInfo struct {
	CardToken    string `json:"CardToken"`
	SecurityCode string `json:"SecurityCode"`
	Brand        string `json:"Brand"`
}

type PaymentRequest struct {
	MerchantOrderId string                `json:"MerchantOrderId"`
	Customer        Customer              `json:"Customer"`
	Payment         PaymentTransactionDTO `json:"Payment"`
}

type PaymentTransactionDTO struct {
	Type           string         `json:"Type"`
	Amount         int            `json:"Amount"`
	Installments   int            `json:"Installments"`
	SoftDescriptor string         `json:"SoftDescriptor"`
	CreditCard     CreditCardInfo `json:"CreditCard"`
}

type Link struct {
	Method string `json:"Method"`
	Rel    string `json:"Rel"`
	Href   string `json:"Href"`
}

type Payment struct {
	ServiceTaxAmount    int            `json:"ServiceTaxAmount"`
	Installments        int            `json:"Installments"`
	Capture             bool           `json:"Capture"`
	Authenticate        bool           `json:"Authenticate"`
	CreditCard          CreditCardInfo `json:"CreditCard"`
	ProofOfSale         string         `json:"ProofOfSale"`
	Tid                 string         `json:"Tid"`
	AuthorizationCode   string         `json:"AuthorizationCode"`
	SoftDescriptor      string         `json:"SoftDescriptor"`
	PaymentId           string         `json:"PaymentId"`
	Type                string         `json:"Type"`
	Amount              int            `json:"Amount"`
	Currency            string         `json:"Currency"`
	Country             string         `json:"Country"`
	ExtraDataCollection []string       `json:"ExtraDataCollection"`
	Status              int            `json:"Status"`
	Links               []Link         `json:"Links"`
}

type TransactionResponse struct {
	MerchantOrderId string   `json:"MerchantOrderId"`
	Customer        Customer `json:"Customer"`
	Payment         Payment  `json:"Payment"`
}
