package models

type TapPayRequest struct {
	Prime      string     `json:"prime"`
	PartnerKey string     `json:"partner_key"`
	MerchantID string     `json:"merchant_id"`
	Amount     int        `json:"amount"`
	Details    string     `json:"details"`
	Cardholder Cardholder `json:"cardholder"`
}

type Cardholder struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	Email       string `json:"email"`
}
