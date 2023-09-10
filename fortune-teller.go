package main

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TOKEN = "XXX"

var bot *tgbotapi.BotAPI
var chatId int64

var fortuneTellerNames = [3]string{"денис", "ден", "денчик"}

var answers = []string{
	"Да",
	"Нет",
	"Не откладывай на завтра то, что можешь сделать сегодня. Марк Твен",
	"Лучше сожалеть о том, что сделал, чем о том, что не сделал.  Джоан Роулинг",
	"Искусство счастья - это умение уметь отличать важное от неважного. Далай Лама",
	"Жизнь слишком коротка, чтобы тратить ее на ненависть и злобу. Далай Лама",
	"Счастье не в том, чтобы иметь то, чего нет, а в том, чтобы ценить то, что есть.  Марк Твен",
	"Чем больше узнаешь, тем меньше знаешь. Льюис Кэрролл",
	"Смелость не отсутствие страха, а способность преодолевать его.  Марк Твен",
	"Не печалься о прошлом, не беспокойся о будущем, сосредоточься на настоящем. Будда",
	"Лучший способ предсказать будущее - создать его. Петер Друкер",
	"Счастье — это не постоянное состояние, это мгновенный выбор. Габриель Гарсиа Маркес",
	"Путешествие тысячи миль начинается с первого шага. Лао-Цзы",
	"Истинная сила проявляется в том, как мы обращаемся с теми, кто слабее нас. Дж.К. Роулинг",
	"Ты отвечаешь за то, что приручил. Антуан де Сент-Экзюпери",
	"Самое важное в жизни - это быть добрым и заботиться о других.Льюис Кэрролл",
	"Судьба не влияет на того, кто не делает ничего. Харпер Ли",
	"Невозможное становится возможным, когда ты веришь в себя. Валериан",
	"Жизнь - это путешествие, а не назначение. Ральф Уолдо Эмерсон",
	"Человек делает свою судьбу своими собственными руками. Фридрих Ницше",
	"Самое сложное в жизни - это принимать решения. Габриэль Гарсиа Маркес",
	"Счастье - это когда ты доволен тем, что у тебя есть. Лев Толстой",
	"Жизнь - это не ждать, когда закончится дождь, а учиться танцевать под дождем. Сенека",
	"Помни, что ты умрешь, и все станет ясно. Стив Джобс",
	"Наилучший способ предсказать будущее - это его создать. Линкольн",
	"Ничто не укрепляет душу так, как забота о других. Карл Густав Юнг",
	"Счастье - это не цель, а способ жизни. Эйнштейн",
	"Судьба человека зависит от его характера. Шопенгауэр",
	"Последний шанс - это всегда тот, который ты упустил. Оруэлл",
	"Следуй своей мечте, и мир откроется перед тобой. Коэльо",
	"Чем больше ты даешь, тем больше получаешь. Карнеги",
	"Жизнь - это не ожидание бури, а учение танцевать под дождем.  Сэнди Шоу",
}

func connectWithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to Telegram ")
	}
}
func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msgConfig)
}

func isMessageForFortuneTeller(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}
	msgInLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range fortuneTellerNames {
		if strings.Contains(msgInLowerCase, name) {
			return true
		}

	}
	return false
}
func getFortuneTellersAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatId, getFortuneTellersAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	connectWithTelegram()
	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatId = update.Message.Chat.ID
			sendMessage("Задай свой вопрос, назвав меня по имени. Ответом на вопрос должны быть \"Да\" либо \"Нет\". Например, \" Денис, я готов сменить работу?\" либо \"Денис, я действительно хочу отправиться на эту вечеринку?\"")
		}
		if isMessageForFortuneTeller(&update) {
			sendAnswer(&update)

		}
	}
}
