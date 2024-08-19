package models

type InitiateMoMoPaymentRequest struct {
	Country     string `json:"country"`
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	Reference   string `json:"reference"`
	DLCode      string `json:"dl_code"`
	BankCode    string `json:"bank_code"`
	AccountNum  string `json:"account_num"`
	AccountName string `json:"account_name"`
	Description string `json:"description"`
	WebhookURL  string `json:"webhook_url"`
}

type InitiateMoMoPaymentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Reference string `json:"reference"`
		Status    string `json:"status"`
	} `json:"data"`
}
