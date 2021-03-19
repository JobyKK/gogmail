package main

import (
	"fmt"
	"gogmail"
	"os"
)

func main()  {
	email := os.Getenv("email")
	if email == "" {
		panic("email is not set")
	}
	password := os.Getenv("password")
	if password == "" {
		panic("password is not set")
	}git
	client := gogmail.NewGmailClient(email, password)

	to := []string{"123@gmail.com"}
	data := gogmail.GmailSendData{Subject: "1", Body: "2"}
	fmt.Println(data)
	client.SendEmail(to, &gogmail.GmailSendData{"Test subject", "Test Body"})
}
