package delivery

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type botFeature interface {
	GetFirstName() string
	SendEarnedToken(userID string, value int32) error
}

type llmFeature interface {
	GenerateHelpText(word string) string
}

type httpHandler struct {
	bot botFeature
	llm llmFeature
}

func NewHttpHandler(bot botFeature, llm llmFeature) httpHandler {
	return httpHandler{bot: bot, llm: llm}
}

func (hh httpHandler) RootPathHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	message := fmt.Sprintf("Hello from %s", hh.bot.GetFirstName())

	err := writeJSON(w, http.StatusOK, envelope{"message": message}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
		return
	}
}

func (hh httpHandler) GenerateHelpTextHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var input struct {
		Word string `json:"word"`
	}

	err := readJSON(r, &input)
	if err != nil {
		serverErrorResponse(w, r, err)
		return
	}

	resp := hh.llm.GenerateHelpText(input.Word)

	err = writeJSON(w, http.StatusOK, envelope{"message": resp}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
		return
	}
}

func (hh httpHandler) SendTgbotToken(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var input struct {
		UserID string `json:"userID"`
		Value  int32  `json:"value"`
	}

	err := readJSON(r, &input)
	if err != nil {
		serverErrorResponse(w, r, err)
		return
	}

	err = hh.bot.SendEarnedToken(input.UserID, input.Value)
	if err != nil {
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"message": "sent telegram token"}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
		return
	}
}
