// routes/check_payment_status.go
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

// CheckPaymentStatusHandler handles requests to check payment status.
func CheckPaymentStatusHandler(c *gin.Context) {
	reference := c.Param("reference")
	url := fmt.Sprintf("%s/external/payment/status/%s", constants.BaseURL, reference)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
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

	var response models.CheckPaymentStatusResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("Error parsing response JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	if response.Success {
		c.JSON(http.StatusOK, gin.H{
			"status": response.Data.Status,

			"description":     response.Data.Description,
			"reference":       response.Data.Reference,
			"clientReference": response.Data.ClientReference,
			"transDate":       response.Data.TransDate,
		})
	} else {
		log.Printf("Status enquiry failed: %s\n", response.Message)
		c.JSON(http.StatusInternalServerError, gin.H{"error": response.Message})
	}
}
