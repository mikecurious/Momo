package models

type MomoMNO struct {
	BankName              string `json:"bank_name"`
	BankCode              string `json:"bank_code"`
	CountryCode           string `json:"country_code"`
	CountryCodeText       string `json:"country_code_text"`
	CountryCurrencySymbol string `json:"country_currency_symbol"`
	Type                  string `json:"type"`
}

type GetMomoMNOsResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []MomoMNO `json:"data"`
}
