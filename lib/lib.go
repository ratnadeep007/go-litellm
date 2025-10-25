package lib

func GenerateText(generationRequest TextGenerationRequest) TextGenerationResponse {
	openai := OpenAI{}
	response := openai.GenerateResponse(generationRequest)
	return response
}
