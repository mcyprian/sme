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

Please review your following information:
ORDER NUMBER: %06d
ORDER DATE: %s

%s

Your return code: %02X

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

func sendOrderMail(recipient string, orderNum uint, orderTime time.Time, returnCode []byte, goods string) {
	/*    sendMail(recipient,
	                   fmt.Sprintf("Your FlyIT order %06d", orderNum)
		               fmt.Sprintf(orderMsg, orderNum, orderTime.Format("2006-01-31"), goods, returnCode)
	*/

	fmt.Println(fmt.Sprintf("Your FlyIT order %06d", orderNum))
	fmt.Println(fmt.Sprintf(orderMsg, orderNum, orderTime.Format("2006-01-31"), goods, returnCode))
}
