package loger

import (
	"log"
	"os"
)

var (
	outfile, _ = os.OpenFile("logs/telegram.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	LogFile    = log.New(outfile, "", 0)
)

// обработка и запись всех ошибок в лог файл
func ForError(err error) {
	if err != nil {
		LogFile.Fatalln(err)
	}
}
