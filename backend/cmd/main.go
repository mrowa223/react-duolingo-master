package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
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

	tgBotToken := os.Getenv("TG_BOT_TOKEN")

	bot, err := gotgbot.NewBot(tgBotToken, &gotgbot.BotOpts{})
	if err != nil {
		log.Fatal(err)
		return
	}

	_ = bot

	router := httprouter.New()

	router.Handle(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte(fmt.Sprintf("Hello from %s", bot.FirstName)))
	})

	address := "localhost:5000"
	log.Printf("starting server at %s...", address)
	err = http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
