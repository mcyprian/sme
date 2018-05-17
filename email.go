package main

import (
	"fmt"
	"net/smtp"
	"time"
)

const (
	SENDER     = "flyitservice@gmail.com"
	PASSWD     = "fly1tmetla"
	GMAIL_SMTP = "smtp.gmail.com"
	PORT       = "587"
)

var orderMsg = `
Thanks for your order!

Please review your order:
ORDER NUMBER: %02X
ORDER DATE: %s

%s

In case you have some questions contact us:
E-mail: flyitservice.gmail.com.
Webpage: http://159.122.183.212:32122

Your FlyIT team.
`

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

func sendOrderMail(recipient string, orderNum []byte, goods string) {
	/*    sendMail(recipient,
	      fmt.Sprintf("Your FlyIT order %s", orderNum),
	      fmt.Sprintf(orderMsg, orderNum, time.Now(), goods))
	*/

	fmt.Println(fmt.Sprintf("Your FlyIT order %02X", orderNum))
	fmt.Println(fmt.Sprintf(orderMsg, orderNum, time.Now().Local().Format("2006-01-31"), goods))
}
