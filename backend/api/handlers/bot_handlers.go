package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/schema"
)

var LLM *googleai.GoogleAI

var Docs *[]schema.Document

func (h *Handler) SendPromptHandler(w http.ResponseWriter, r *http.Request) {
	*Docs = append(*Docs, schema.Document{PageContent: fmt.Sprintf(`Today's date: %v`, time.Now().Format("02-01-2006"))})
	type promptReqBody struct {
		Prompt string `json:"prompt"`
	}
	req := promptReqBody{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to read prompt request body",
			"error_dets": err.Error(),
		}, http.StatusBadRequest)
		return
	}

	globalChain := chains.LoadStuffQA(LLM)
	answer, err := chains.Call(context.Background(), globalChain, map[string]any{
		"input_documents": *Docs,
		"question":        req.Prompt,
	}, chains.WithTemperature(0.5))
	if err != nil {
		WriteJSON(w, map[string]string{
			"message":    "Failed to generate bot response",
			"error_dets": err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	refAnswer := strings.ReplaceAll(answer["text"].(string), "\n", "<br>")
	WriteJSON(w, map[string]string{
		"response": refAnswer,
	}, http.StatusOK)
}
