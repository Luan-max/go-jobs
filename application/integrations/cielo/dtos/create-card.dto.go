package dtos

type CreditCardDto struct {
	CustomerName   string `json:"CustomerName"`
	CardNumber     string `json:"CardNumber"`
	Holder         string `json:"Holder"`
	ExpirationDate string `json:"ExpirationDate"`
	Brand          string `json:"Brand"`
}

type CardAPIResponse struct {
	CardToken string `json:"CardToken"`
	Links     struct {
		Method string `json:"Method"`
		Rel    string `json:"Rel"`
		Href   string `json:"Href"`
	} `json:"Links"`
}
