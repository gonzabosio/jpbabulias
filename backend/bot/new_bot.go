package bot

import (
	"context"
	"fmt"
	"os"

	"github.com/gonzabosio/jpbabulias/api/handlers"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/schema"
)

func NewLLM() error {
	ctx := context.Background()
	api_key := os.Getenv("GEMINI_API_KEY")
	llm, err := googleai.New(ctx, googleai.WithAPIKey(api_key))
	if err != nil {
		return fmt.Errorf("Failed to run gemini model: %v", err)
	}

	docs := []schema.Document{
		{
			PageContent: `You are the virtual assistant of the dentist Juan Pablo Babulias. Your main function is to give information or guide the user through the website`,
			Metadata:    map[string]interface{}{"topic": "your identity"},
		},
		{
			PageContent: `Respond in Spanish only, and do not use any English words in your answers. Do not mention that you are consulting documents or providing information from documents. 
			If you receive a greeting, say hello too, be kind and offer help, don't give specific information.`,
			Metadata: map[string]interface{}{"topic": "response considerations"},
		},
		{
			PageContent: `
			To check the schedule and make an appointment, go to "Pedir Turno" section of the website. If someone asks you to schedule an appointment through you, you can't.
			To see the treatments performed and their results, visit the 'Tratamientos' section, which showcases the doctor's experience and success stories.
			`,
			Metadata: map[string]interface{}{"topic": "website navigation"},
		},
		{
			PageContent: `The office address is Avenida Juan de Garay 2055. In San Francisco, Córdoba, Argentina`,
			Metadata:    map[string]interface{}{"topic": "office address"},
		},
		{
			PageContent: `The opening hours are: 
			- Monday to Thursday: 8:00 to 12:00 and 16:00 to 20:00. 
			- Friday: 8:00 to 16:00.`,
			Metadata: map[string]interface{}{"topic": "opening hours"},
		},
		{
			PageContent: `Doctor's social media: 
			- Instagram: od.jpbabulias
			- Whatsapp: +54-3564-590331`,
			Metadata: map[string]interface{}{"topic": "doctor's social media"},
		},
	}

	handlers.LLM = llm
	handlers.Docs = &docs

	return nil
}
