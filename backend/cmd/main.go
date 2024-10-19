package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mrowa223/react-hackaton/backend/internal/delivery"
	"github.com/mrowa223/react-hackaton/backend/internal/feature"
)

func main() {
	var envFilePath string

	flag.StringVar(&envFilePath, "env-path", ".env", ".env file path (only for development environment)")
	flag.Parse()

	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("loaded .env file")

	bot := feature.NewTelegramBot(os.Getenv("TG_BOT_TOKEN"))
	llm := feature.NewLLMFeature(os.Getenv("GEMINI_API_KEY"))

	httpHandler := delivery.NewHttpHandler(bot, llm)
	router := delivery.NewHttpRouter(httpHandler)

	go bot.Start()

	address := "localhost:5001"
	log.Printf("[HTTP] starting server at %s...", address)
	err = http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
