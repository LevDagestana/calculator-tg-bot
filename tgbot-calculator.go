package main

import (
	"strconv"

	tgApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgApi.BotAPI

const TOKEN = "7021454736:AAHURCxd4jpmCwLacSfq8VT_-zJJ3u9qZLE"

func main() {
	var err error
	bot, err = tgApi.NewBotAPI(TOKEN)
	if err != nil {
		panic("Cannod created bot")
	}

	u := tgApi.NewUpdate(0)

	updates := bot.GetUpdatesChan(u)
	var values = []float64{}
	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == "/start" {
				msg := tgApi.NewMessage(update.Message.Chat.ID, "Это бот проводит четыре операции \n(сумма, разность, произведение, деление) \nс двумя цифрами.")
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

				msgNumFirst := tgApi.NewMessage(update.Message.Chat.ID, "Введите первое число")
				bot.Send(msgNumFirst)
				values = []float64{}
				continue
			}
			isNum := func(s string) bool {
				_, err := strconv.ParseFloat(s, 64)
				return err == nil
			}(update.Message.Text)

			if !isNum {
				msg := tgApi.NewMessage(update.Message.Chat.ID, "Вы ввели не число")
				bot.Send(msg)
				continue
			}

			if len(values) == 0 {
				text, _ := strconv.ParseFloat(update.Message.Text, 32)
				values = append(values, float64(text))
				msgNumFirst := tgApi.NewMessage(update.Message.Chat.ID, "Введите второе число")
				bot.Send(msgNumFirst)
				continue
			}
			if len(values) == 1 {
				text, _ := strconv.ParseFloat(update.Message.Text, 32)
				values = append(values, text)
				sum := strconv.FormatFloat(values[0]+values[1], 'g', 10, 64)
				dif := strconv.FormatFloat(values[0]-values[1], 'g', 10, 64)
				mul := strconv.FormatFloat(values[0]*values[1], 'g', 10, 64)
				div := strconv.FormatFloat(values[0]/values[1], 'g', 10, 64)

				message := "Сумма = " + sum + "\nРазница = " + dif + "\nПроизведение = " + mul + "\nЧастное = " + div

				msgNumSecond := tgApi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(msgNumSecond)
			}

		}

	}

}
