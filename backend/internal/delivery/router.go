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

	return recoverPanic(logMiddleware(cors.AllowAll().Handler(router)))
}
