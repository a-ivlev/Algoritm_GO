package notification

import (
	"fmt"
	"net/smtp"
	"shop/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Notification interface {
	SendOrderCreated(order *models.Order) error
}

type telegramBot struct {
	chatID int64
	tgBot  *tgbotapi.BotAPI
}

func (s *telegramBot) SendOrderCreated(order *models.Order) error {
	text := fmt.Sprintf("new order %d\n\nname: %s\nphone: %s\n\nemail: %s\n",
		order.ID, order.CustomerName, order.CustomerPhone, order.CustomerEmail)

	msg := tgbotapi.NewMessage(s.chatID, text)

	_, err := s.tgBot.Send(msg)
	return err
}

func NewTelegramBot(token string, chatID int64) (Notification, error) {
	tgBot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &telegramBot{
		chatID: chatID,
		tgBot:  tgBot,
	}, nil
}

type SendEmailNotif interface {
	SendEmailOrderNotification(order *models.Order) error
}

type sendEmail struct {
	login string
	password string
}

func (m *sendEmail) SendEmailOrderNotification(order *models.Order) error {
	// server we are authorized to send email through
	host := "mail.ru"//m.host
	// user we are authorizing as
	from := m.login
	// use we are sending email to
	to := order.CustomerEmail
	// Create the authentication for the SendMail()
	// using PlainText, but other authentication methods are encouraged
	auth := smtp.PlainAuth("", from, m.password, host)
	// NOTE: Using the backtick here ` works like a heredoc, which is why all the
	// rest of the lines are forced to the beginning of the line, otherwise the
	// formatting is wrong for the RFC 822 style
	message := `To: "Some User" <someuser@example.com>
From: "Other User" <otheruser@example.com>
Subject: Testing Email From Go!!

This is the message we are sending. That's it!
`
	if err := smtp.SendMail(host+":25", auth, from, []string{to}, []byte(message)); err != nil {
		fmt.Println("Error SendMail: ", err)
		return err
	}
	fmt.Println("Email Sent!")
	return nil
}

func NewEmail(loginEmail string, passwordEmail string) (SendEmailNotif, error) {
	return &sendEmail{
		login: loginEmail,
		password: passwordEmail,
	}, nil
}
