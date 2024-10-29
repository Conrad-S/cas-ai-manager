package ai_orchestration_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RequestBody struct {
	Model      string   `json:"model"`
	Version    string   `json:"version"`
	Timeout    int      `json:"timeout"`
	Choices    []Choice `json:"choices"`
	CallOpenAI bool     `json:"callOpenAI"`
}

type Response struct {
	Model    string `json:"model"`
	Version  string `json:"version"`
	Region   string `json:"region"`
	Response string `json:"response"`
	Status   string `json:"status"`
}

type Choice struct {
	Model   string `json:"model"`
	Version string `json:"version"`
	Region  string `json:"region"`
}

func CallAzureOpenAI(url string, body []byte, timeout time.Duration) (string, error) {
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if errorMsg, ok := result["error"]; ok {
		return "", fmt.Errorf("error from API: %v", errorMsg)
	}

	respBody, err := json.Marshal(result)
	if err != nil {
		return "", fmt.Errorf("failed to marshal response: %w", err)
	}
	return string(respBody), nil
}
