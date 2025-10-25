package lib

type TextGenerationRequest struct {
	Model             string
	Prompt            string
	System            string
	Max_tokens        int
	Temperature       float32
	Top_p             float32
	Frequency_penalty float32
	Presence_penalty  float32
}

type OpenAI struct {
	apiKey  string
	baseURL string
}

type TextGenerationContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Role string `json:"role"`
}

type TextGenerationOutput struct {
	ID      string                  `json:"id"`
	Type    string                  `json:"type"`
	Status  string                  `json:"status"`
	Content []TextGenerationContent `json:"content"`
}

type TextGenerationResponse struct {
	ID     string                 `json:"id"`
	Output []TextGenerationOutput `json:"output"`
}
