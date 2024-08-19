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

// GenerateOTPHandler handles requests to generate OTP.
func GenerateOTPHandler(c *gin.Context) {
	var requestBody models.GenerateOTPRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	url := fmt.Sprintf("%s/otp/generate/dynamic-link", constants.ESAURL)

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

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response body"})
		return
	}

	// Parse response JSON
	var response models.GenerateOTPResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("Error parsing response JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response JSON"})
		return
	}

	// Check if OTP link generation was successful
	if response.Data.Status {
		c.JSON(http.StatusOK, gin.H{
			"message": "OTP generated successfully",
			"otp_sid": response.Data.OTPSID,
			"status":  response.Data.Status,
		})
	} else {
		log.Printf("OTP generation failed: %v\n", response.Data.Message)
		c.JSON(http.StatusInternalServerError, gin.H{"error": response.Data.Message})
	}
}
