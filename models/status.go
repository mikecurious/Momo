package models

type CheckPaymentStatusResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Status          string `json:"status"`
		Description     string `json:"description"`
		Reference       string `json:"reference"`
		ClientReference string `json:"clientReference"`
		TransDate       string `json:"transDate"`
	} `json:"data"`
}
