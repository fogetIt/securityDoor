package utils


import (
	"fmt"
	"net/smtp"
	"strings"
)


func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte(fmt.Sprintf(
		"To:%s\r\nFrom%s<%s>\r\nSubject: %s\r\n%s\r\n\r\n%s",
		to, user, user, subject, contentType, body))
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}


func SendEmail(Code int) bool {
	return true
}


func SendMessage(Code int) bool {
	return true
}

