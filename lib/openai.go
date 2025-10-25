package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// GenerateResponse takes a TextGenerationRequest and returns a generated response string from OpenAI.
func (o *OpenAI) GenerateResponse(generationRequest TextGenerationRequest) TextGenerationResponse {
	o.apiKey = checkEnvVar("OPENAI_API_KEY")
	if o.apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set")
		return TextGenerationResponse{}
	}

	if o.baseURL == "" {
		o.baseURL = "https://api.openai.com/v1/responses"
	}

	client := &http.Client{}

	payload := map[string]any{
		"model": generationRequest.Model,
		"input": generationRequest.Prompt,
		// "max_tokens":        generationRequest.Max_tokens,
		// "temperature":       generationRequest.Temperature,
		// "top_p":             generationRequest.Top_p,
		// "frequency_penalty": generationRequest.Frequency_penalty,
		// "presence_penalty":  generationRequest.Presence_penalty,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling payload:", err)
		os.Exit(1)
	}

	req, err := http.NewRequest("POST", o.baseURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	// Add the API key to the request headers
	req.Header.Add("Authorization", "Bearer "+o.apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making API request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	var textGenerationResponse TextGenerationResponse
	err = json.Unmarshal(body, &textGenerationResponse)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		os.Exit(1)
	}

	return textGenerationResponse
}
