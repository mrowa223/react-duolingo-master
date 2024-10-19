package feature

import (
	"fmt"
	"log"

	tg "github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

type telegramBotFeature struct {
	bot *tg.Bot
}

func NewTelegramBot(token string) *telegramBotFeature {
	bot, err := tg.NewBot(token, &tg.BotOpts{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &telegramBotFeature{bot: bot}
}

func (f telegramBotFeature) GetFirstName() string {
	return f.bot.FirstName
}

func (f telegramBotFeature) Start() {
	dispatcher := ext.NewDispatcher(nil)
	updater := ext.NewUpdater(dispatcher, nil)

	dispatcher.AddHandler(handlers.NewCommand("start", handleStart))

	err := updater.StartPolling(f.bot, nil)
	if err != nil {
		log.Fatalf("Failed to start polling: %v", err)
	}

	log.Printf("[TGBOT] %s has been started...\n", f.bot.User.Username)

	updater.Idle()
}

func handleStart(b *tg.Bot, ctx *ext.Context) error {
	chatID := ctx.EffectiveChat.Id

	inlineKeyboard := tg.InlineKeyboardMarkup{
		InlineKeyboard: [][]tg.InlineKeyboardButton{
			{
				{
					Text: "Open Web App",
					WebApp: &tg.WebAppInfo{
						Url: "https://hackaton-bnb.vercel.app/",
					},
				},
			},
		},
	}

	_, err := b.SendMessage(chatID, "Welcome! Click the button below to open the Web App:", &tg.SendMessageOpts{
		ReplyMarkup: inlineKeyboard,
	})
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (f telegramBotFeature) SendEarnedToken(userID string, value int32) error {
	panic("not implemented")
	return nil
}
