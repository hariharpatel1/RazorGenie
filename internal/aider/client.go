package aider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	aiderAPIEndpoint = "https://habibiswitch.openai.azure.com/openai/deployments/Habibi-4o/chat/completions?api-version=2024-02-15-preview" // Example endpoint
	apiKey           = "96a66b32cf8e424c84632e221bbec779"                                                                                   // Replace with your Aider API key
)

// AiderClient represents a client for interacting with Aider's API.
type AiderClient struct {
	apiKey string
	client *http.Client
}

// NewAiderClient creates a new instance of AiderClient.
func NewClient(apiKey string) *AiderClient {
	return &AiderClient{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

// GenerateCode uses Aider's API to generate code based on the prompt.
func (c *AiderClient) GenerateCode(prompt string) (string, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"prompt":      prompt,
		"max_tokens":  1000,
		"temperature": 0.7,
		"top_p":       1,
	})
	if err != nil {
		return "", fmt.Errorf("error creating request body: %v", err)
	}

	req, err := http.NewRequest("POST", aiderAPIEndpoint, bytes.NewReader(requestBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to Aider API: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("error decoding Aider API response: %v", err)
	}

	fmt.Println("[Result]: %v", result)

	if result["error"] != nil {
		return "", fmt.Errorf("aider api error: %v", result["error"])
	}

	return result["choices"].([]interface{})[0].(map[string]interface{})["text"].(string), nil
}
