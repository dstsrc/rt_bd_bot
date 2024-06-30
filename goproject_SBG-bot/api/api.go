package api

import (
	//	"encoding/csv"
	//	"fmt"
	//	"goproject_SBG-bot/service"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type service interface {
	Chek_avtorisation(chatID int64) bool
	Out_list(chatID int64) string
	AddName(chatID int64) error
	AddNameWork(text_vvod string, chatID int64) error
	DeleteName(chatID int64) (string, error)
	DeleteNameWork(text_vvod string, chatID int64) error
	Сancel(chatID int64) string
	Get_worker() ([][]int64, []string)
	Get_previous(chatID int64) string
	EnterName(text string, chatID int64) error
	EnterDate(text string, chatID int64) error
}

func Run(s service) {

	bot, err := tgbotapi.NewBotAPI("6825004600:AAE-h0KLA3tJ4AjIyxH6PU0cCVIrPDFZrsM")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	go worker(bot, s)
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // Есть новое сообщение
			text := update.Message.Text      // Текст сообщения
			chatID := update.Message.Chat.ID //  ID чата
			userID := update.Message.From.ID // ID пользователя
			var replyMsg string

			log.Printf("[%s](%d) %s", update.Message.From.UserName, userID, text)

			// Анализируем текст сообщения и записываем ответ в переменную

			replyMsg = Distribution_answers(s, text, chatID)

			// Отправляем ответ
			msg := tgbotapi.NewMessage(chatID, replyMsg) // Создаем новое сообщение
			//			msg.ReplyToMessageID = update.Message.MessageID // Указываем сообщение, на которое нужно ответить

			bot.Send(msg)
		}
	}
}

func worker(bot *tgbotapi.BotAPI, s service) {

	for {
		slice_ID, to_name := s.Get_worker()

		for i, name := range to_name {
			for _, v_ID := range slice_ID[i] {
				replyMsg := "скоро ДР у " + name
				msg := tgbotapi.NewMessage(v_ID, replyMsg)
				bot.Send(msg)
			}
		}
		time.Sleep(time.Second * 10)
	}
}

func Distribution_answers(s service, text_vvod string, chatID int64) string {
	var replyMsg string

	if text_vvod == "/cancel" {
		//		fmt.Println("++Cancel++")
		replyMsg = s.Сancel(chatID)
		return replyMsg
	}
	if s.Chek_avtorisation(chatID) == false {
		replyMsg = Autorisation(s, text_vvod, chatID)
	} else {

		if text_vvod == "/list_name" {
			replyMsg = s.Out_list(chatID)

		} else if text_vvod == "/add_name" {
			replyMsg = AddName(s, chatID)

		} else if text_vvod == "/delete_name" {
			replyMsg = DeleteName(s, chatID)

		} else if s.Get_previous(chatID) == "Add" {
			replyMsg = AddNameWork(s, text_vvod, chatID)

		} else if s.Get_previous(chatID) == "Delete" {
			replyMsg = DeleteNameWork(s, text_vvod, chatID)

		} else {
			replyMsg = Menu()
		}
	}

	return replyMsg
}

func Autorisation(s service, text_vvod string, chatID int64) string {

	err := s.EnterName(text_vvod, chatID)
	if err != nil {
		return "Авторизация\n\nИмя пользователя не введео или введено неверно\nВведите имя пользователя\nили отмена действия /cancel"
	}

	err = s.EnterDate(text_vvod, chatID)
	if err != nil {
		return "Авторизация\n\nДата рождения пользователя не введена или введена неверно\nВведите дату рождения пользователя в формате:\nГГГГ-ММ-ДД\nнапример:\n2003-10-28\nили отмена действия /cancel"
	}

	return "Авторизация завершена" + "\n" + "\n" + Menu()
}

func AddName(s service, chatID int64) string {
	err := s.AddName(chatID)
	if err != nil {
		return "Поользватель не прошел регистрацию"
	}
	return "Введите имя пользователя\nили отмена действия /cancel"
}

func AddNameWork(s service, text_vvod string, chatID int64) string {
	err := s.AddNameWork(text_vvod, chatID)
	if err != nil {
		return "Пользователь не существует\n" +
			"Введите другое имя\n" +
			"или отмена действия /cancel"
	}
	return "Пользователь успешно добавлен"
}

func DeleteName(s service, chatID int64) string {
	listName, err := s.DeleteName(chatID)
	if err != nil {
		return "Список на кого подписан пуст"
	}
	return "Введите имя пользователя из списка подписанных:" + listName + "\nили отмена действия /cancel"
}

func DeleteNameWork(s service, text_vvod string, chatID int64) string {
	err := s.DeleteNameWork(text_vvod, chatID)
	if err != nil {
		return "Пользователь с указанным именем не зарегестрирован либо\n" +
			"Вы не подписаны на пользователя с указанным именем\n" +
			"Введите другое имя\n" +
			"или отмена действия /cancel"
	}
	return "Пользователь успешно удален"
}

func Menu() string {
	text := "Доступны следующие действия:\n" +
		"получить список пользователей - /list_name\n" +
		"подписаться на пользователя - /add_name\n" +
		"отписаться от пользователя - /delete_name\n"

	return text
}
