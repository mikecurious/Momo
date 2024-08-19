package models

type ValidateAccountResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Name     string `json:"name"`
		Accounts []struct {
			AccountID string `json:"account_id"`
			Currency  string `json:"currency"`
		} `json:"accounts"`
	} `json:"data"`
}
