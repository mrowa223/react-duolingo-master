package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mrowa223/react-hackaton/backend/internal/delivery"
	"github.com/mrowa223/react-hackaton/backend/internal/feature"
	"github.com/rs/cors" // Import the CORS package
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

	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow specific methods
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // If you need credentials (cookies, etc.)
	})

	address := "localhost:5001"
	log.Printf("[HTTP] starting server at %s...", address)
	err = http.ListenAndServe(address, corsHandler.Handler(router)) // Wrap the router with the CORS handler
	if err != nil {
		log.Fatal(err)
		return
	}
}
