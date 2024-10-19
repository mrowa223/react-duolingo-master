package delivery

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func NewHttpRouter(httpHandler httpHandler) http.Handler {
	router := httprouter.New()

	router.Handle(http.MethodGet, "/", httpHandler.RootPathHandler)
	router.Handle(http.MethodPost, "/v1/llm/help-text", httpHandler.GenerateHelpTextHandler)
	// router.Handle(http.MethodPost, "/v1/tgbot/send-token", httpHandler.SendTgbotToken)

	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5001"},                   // Allow frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow specific methods
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // If you need credentials (cookies, etc.)
	})

	return recoverPanic(logMiddleware(corsHandler.Handler(router)))
}
