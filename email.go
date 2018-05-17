package main

import (
	"fmt"
	"net/smtp"
)

const (
	SENDER     = "flyitservice@gmail.com"
	PASSWD     = "fly1tmetla"
	GMAIL_SMTP = "smtp.gmail.com"
	PORT       = "587"
)

func sendMail(recipient string, subject string, body string) {
	msg := "From: " + SENDER + "\n" +
		"To: " + recipient + "\n" +
		"Subject: " + subject +
		body

	err := smtp.SendMail(fmt.Sprintf("%s:%s", GMAIL_SMTP, PORT),
		smtp.PlainAuth("", SENDER, PASSWD, GMAIL_SMTP),
		SENDER, []string{recipient}, []byte(msg))

	if err != nil {
		panic(err)
	}
}
