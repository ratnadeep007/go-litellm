package handlers

import (
	"go-litellm/lib"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func generateResponse(prompt string, model string) lib.TextGenerationResponse {
	generateRequest := lib.TextGenerationRequest{
		Model:  model,
		Prompt: prompt,
	}
	return lib.GenerateText(generateRequest)
}

type messageRequest struct {
	Message string `json:"message"`
	Model   string `json:"model"`
}

func GenerateResponse(c echo.Context) error {
	log.Info("Generating response")
	var req messageRequest

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	if req.Message == "" {
		return c.String(http.StatusBadRequest, "Message is required")
	}
	prompt := req.Message
	model := req.Model

	if prompt == "" {
		return c.String(http.StatusBadRequest, "Message is required")
	}

	if model == "" {
		return c.String(http.StatusBadRequest, "Model is required")
	}

	response := generateResponse(prompt, model)
	return c.JSON(http.StatusOK, response)
}
