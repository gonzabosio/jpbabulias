package main

import (
	"log"
	"net/http"
	"os"

	backend "github.com/gonzabosio/jpbabulias/api"
	"github.com/gonzabosio/jpbabulias/bot"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../backend/.env"); err != nil {
		log.Printf("environment variables loading error: %v\n", err)
	}
	if err := bot.NewLLM(); err != nil {
		log.Fatal(err.Error())
	}
	r, err := backend.NewRouter()
	if err != nil {
		log.Fatalf("Unable to setup backend server: %v", err)
	}
	if err := http.ListenAndServe(":"+os.Getenv("BACKEND_PORT"), r); err != nil {
		log.Fatalf("Unable to start backend server on port %s: %v", os.Getenv("BACKEND_PORT"), err)
	}
}
