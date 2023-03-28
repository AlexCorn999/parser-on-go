package telegram

import (
	"bytes"
	"fmt"
	"net/http"
	"parser/loger"
)

const (
	chatId       = // your chatID
	BOT_TOKEN    = // your botToken
	TELEGRAM_URL = "https://api.telegram.org/bot"
)

func SendMessage(text string) {
	textAll := fmt.Sprintf("%s", text)
	data := []byte(fmt.Sprintf(`{"chat_id":%d, "text" : "%s", "parse_mode": "HTML"}`, chatId, textAll))
	tx := bytes.NewReader(data)
	_, err := http.Post(fmt.Sprintf("%s%s/sendMessage", TELEGRAM_URL, BOT_TOKEN), "application/json", tx)
	loger.ForError(err)
}
