// sms/bulkSmsBd/bulk_sms_bd.go

package bulkSmsBd

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type BulkSmsBdProvider struct {
	APIKey   string
	SenderID string
}

type singleSMSRequest struct {
	APIKey   string `json:"api_key"`
	SenderID string `json:"senderid"`
	Number   string `json:"number"`
	Message  string `json:"message"`
}

type apiResponse struct {
	ResponseCode   int    `json:"response_code"`
	MessageID      int    `json:"message_id"`
	SuccessMessage string `json:"success_message"`
	ErrorMessage   string `json:"error_message"`
}

func (b BulkSmsBdProvider) Send(to, message string) (string, error) {
	requestData := singleSMSRequest{
		APIKey:   b.APIKey,
		SenderID: b.SenderID,
		Number:   to,
		Message:  message,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://bulksmsbd.net/api/smsapi", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var response apiResponse
	json.Unmarshal(body, &response)

	if response.ResponseCode == 202 {
		return string(body), nil
	} else {
		return string(body), errors.New(response.ErrorMessage)
	}
}
