package tool

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendToMail(form, to, subject, body string) error {
	mailContent := fmt.Sprintf("To:%v\r\n"+
		"Subject: %v\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n\r\n "+
		"%v", to, subject, body)
	message := []byte(mailContent)

	host := "smtp.gmail.com:587"
	user := "quqinyuni@gmail.com"
	password := "bojtmammspmngapa"
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	send_to := strings.Split(to, ";")
	return smtp.SendMail(host, auth, form, send_to, message)
}
