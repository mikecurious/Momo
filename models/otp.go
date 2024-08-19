package models

type GenerateOTPRequest struct {
	ConfigPreset      string                       `json:"config_preset"`
	TransactionParams GenerateOTPTransactionParams `json:"transaction_params"`
}

type GenerateOTPTransactionParams struct {
	SrcAmount string `json:"src_amount"`
	DesAmount string `json:"des_amount"`
	PayeeID   string `json:"payeeId"`
	PayerID   string `json:"payerId"`
}

type GenerateOTPResponse struct {
	Status bool `json:"status"`
	Data   struct {
		Status  bool   `json:"status"`
		OTPSID  string `json:"otp_sid"`
		Message string `json:"message"`
	} `json:"data"`
}
