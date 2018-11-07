package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"strings"
)

// init Loading Config
var (
	HOST, SMTPADDR, USER, PASS string
	PORT                       int
)

type Config struct {
	Dev Env `json:"dev"`
	Pro Env `json:"pro"`
}

type Env struct {
	Email Email `json:"email"`
}

type Email struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	User     string `json:"user"`
	Port     int    `json:"port"`
}

func init() {
	config := Config{}
	filename := "email/config.json"
	fd, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Don't fund Config file %s", filename)
	}
	if err := json.Unmarshal(fd, &config); err != nil {
		log.Fatalf("Load Config Faild %s", err)
	} else {
		HOST = config.Dev.Email.Host
		PORT = config.Dev.Email.Port
		SMTPADDR = fmt.Sprintf("%s:%d", HOST, PORT)
		USER = config.Dev.Email.User
		PASS = config.Dev.Email.Password
	}
}

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
