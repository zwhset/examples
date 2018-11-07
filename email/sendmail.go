package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

// Loading Config
var (
	HOST     = "smtp.163.com"
	PORT     = 25
	SMTPADDR = fmt.Sprintf("%s:%d", HOST, PORT)
	USER     = "user@163.com"
	PASS     = "password"
)

func main() {
	toUser := []string{"user1@qq.com", "user2@qq.com"}
	title := "Re明天出去玩"
	body := "好的"

	err := SendMail(toUser, title, body)
	if err != nil {
		log.Printf("Smtp Sendmail Faild, %s\n", err)
	} else {
		log.Printf("SendMail Success.\n")
	}

}

func SendMail(toUser []string, title, body string) error {
	auth := smtp.PlainAuth("", USER, PASS, HOST)
	msg := makeMsg(toUser, body, title)

	return smtp.SendMail(SMTPADDR, auth, USER, toUser, msg)

}

// make mail message for touser title and body
func makeMsg(toUser []string, body, title string) []byte {

	// more user
	uBody := strings.Join(toUser, ";")

	msg := "From: " + USER + "\r\n" +
		"To: " + uBody + "\r\n" +
		"Subject: " + title + "\r\n\n" +
		body

	return []byte(msg)
}
