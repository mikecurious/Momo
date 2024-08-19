package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"E-transact/constants"
	"E-transact/models"

	"github.com/gin-gonic/gin"
)

// Initiate
func InitiateMoMoPaymentHandler(c *gin.Context) {
	var requestBody models.InitiateMoMoPaymentRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	url := fmt.Sprintf("%s/external/payment/momo", constants.BaseURL)

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error encoding JSON"})
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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

	var response models.InitiateMoMoPaymentResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("Error parsing response JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	if response.Success {
		c.JSON(http.StatusOK, gin.H{
			"message":   "MoMo payment initiated successfully",
			"reference": response.Data.Reference,
			"status":    response.Data.Status,
		})
	} else {
		log.Printf("MoMo payment initiation failed: %s\n", response.Message)
		c.JSON(http.StatusInternalServerError, gin.H{"error": response.Message})
	}
}
