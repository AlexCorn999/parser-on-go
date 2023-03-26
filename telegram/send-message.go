package telegram

import (
	"bytes"
	"fmt"
	"net/http"
	"parser/loger"
)

const (
	chatId       = 1523299067
	BOT_TOKEN    = "5729925133:AAEgZdT5-F8XVfz76mZItVKgBJzIkyLMQQ0"
	TELEGRAM_URL = "https://api.telegram.org/bot"
)

type BotSendMessageID struct {
	Result struct {
		Message_id int
	}
}

func SendMessage(text string) {
	//textAll := fmt.Sprintf("%s", text)
	data := []byte(fmt.Sprintf(`{"chat_id":%d, "text" : "%s", "parse_mod": "HTML", "disable_web_page_preview"}`, chatId, text))
	tx := bytes.NewReader(data)
	_, err := http.Post(fmt.Sprintf("%s%s/sendMessage", TELEGRAM_URL, BOT_TOKEN), "application/json", tx)
	loger.ForError(err)
}
