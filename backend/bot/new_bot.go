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
	llm, err := googleai.New(ctx,
		googleai.WithAPIKey(api_key),
		googleai.WithDefaultModel("gemini-2.0-flash"),
		googleai.WithDefaultTemperature(0.5),
	)
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
			Be kind. Do not greet if you are not greeted first.`,
			Metadata: map[string]interface{}{"topic": "response considerations"},
		},
		{
			PageContent: `
			To check the schedule and make an appointment, go to "Pedir Turno" section of the website. If someone asks you to schedule an appointment through you, you can't.
			To see the treatments performed and their results, visit the 'Tratamientos' section, which showcases the doctor's experience and success stories.
			Login and signup (top-right corner in landing view) to allow the user to make an appointment and manage profile and patient data.`,
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
		{
			PageContent: `When user is logged in, he will have a profile section where he can visualize his appointments and the ones of his saved patients whose data
			can be modified or deleted, but also add new patients to the account. Also they can close the user session, only say it if the ask for it specifically.`,
		},
	}

	handlers.LLM = llm
	handlers.Docs = &docs

	return nil
}
