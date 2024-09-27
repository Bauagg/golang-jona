package utils

import (
	"backend-jona-golang/config"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type BankTransferPayload struct {
	PaymentType        string `json:"payment_type"`
	TransactionDetails struct {
		OrderID     string `json:"order_id"`
		GrossAmount uint64 `json:"gross_amount"`
		ExpireTime  string `json:"expire_time"`
	} `json:"transaction_details"`
	BankTransfer struct {
		Bank string `json:"bank"`
	} `json:"bank_transfer"`
}

type BankTransferResponse struct {
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	TransactionID     string `json:"transaction_id"`
	OrderID           string `json:"order_id"`
	MerchantID        string `json:"merchant_id"`
	GrossAmount       string `json:"gross_amount"`
	Currency          string `json:"currency"`
	PaymentType       string `json:"payment_type"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	VANumbers         []struct {
		Bank     string `json:"bank"`
		VANumber string `json:"va_number"`
	} `json:"va_numbers"`
	FraudStatus string `json:"fraud_status"`
}

func VaNumberBank(data BankTransferPayload) (BankTransferResponse, error) {
	client := resty.New()

	// Encode authorization header
	merchantID := config.SERVER_KEY_MIDTRANS // Ganti dengan merchant ID yang benar
	authorization := "Basic " + base64.StdEncoding.EncodeToString([]byte(merchantID))

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", authorization).
		SetBody(data).
		Post(config.URL_MIDTRANS)

	if err != nil {
		return BankTransferResponse{}, fmt.Errorf("error making POST request: %w", err)
	}

	if resp.StatusCode() != 200 {
		return BankTransferResponse{}, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	var response BankTransferResponse
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return BankTransferResponse{}, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return response, nil
}
