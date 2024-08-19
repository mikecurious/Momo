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

func ValidateMomoAccountHandler(c *gin.Context) {
	countryCode := c.Param("countryCode")
	accountNo := c.Param("accountNo")
	bankCode := c.Param("bankCode")

	url := fmt.Sprintf("%s/external/validate-account/%s/%s/%s", constants.BaseURL, countryCode, accountNo, bankCode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	req.Header.Set("X-MERCHANT-ID", constants.MerchantID)
	req.Header.Set("X-API-KEY", constants.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error making request"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response body"})
		return
	}

	fmt.Printf("HTTP Status Code: %d\n", resp.StatusCode)

	fmt.Printf("Raw Response Body: %s\n", string(body))

	var response models.ValidateAccountResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("Error parsing response JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	if !response.Success {
		log.Printf("API request failed: %v\nError: %v", response.Message, response.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": response.Message})
		return
	}

	c.JSON(http.StatusOK, response)
}
