package service

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendMail(email string) {
	fmt.Println("email",email)
	m := gomail.NewMessage()
	m.SetHeader("From", "dfinity_cycle@163.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Cycle Balance")
	m.SetBody("text/html", " <b>Your Cycle Balance Is Insufficient!</b>")

	d := gomail.NewDialer("smtp.163.com", 465, "dfinity_cycle@163.com", "BOIPFJOEZXZXMGDR")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("send successfully")
}




