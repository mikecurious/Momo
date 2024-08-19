package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"E-transact/constants"
	"E-transact/models"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/momo-mnos/:countryCode", GetMomoMNOsHandler)
	r.GET("/validate-account/:countryCode/:accountNo/:bankCode", ValidateMomoAccountHandler)
	r.POST("/otp/generate", GenerateOTPHandler)
	r.POST("/payment/momo", InitiateMoMoPaymentHandler)
	r.GET("/payment/status/:reference", CheckPaymentStatusHandler)
}

func GetMomoMNOsHandler(c *gin.Context) {
	countryCode := c.Param("countryCode")
	url := fmt.Sprintf("%s/external/momo-mnos/%s", constants.BaseURL, countryCode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("X-MERCHANT-ID", constants.MerchantID)
	req.Header.Set("X-API-KEY", constants.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Printf("HTTP Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Raw Response Body: %s\n", string(body))

	var response models.GetMomoMNOsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error parsing response JSON: %v", err)
	}

	if !response.Success {
		log.Fatalf("API request failed: %v", response.Message)
	}

	if len(response.Data) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No Momo MNOs found for the specified country code."})
		return
	}

	c.JSON(http.StatusOK, response)
}
