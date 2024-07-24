package main

import (
	"gomail/config"
	"log"

	"gopkg.in/gomail.v2"
)

func main() {

	emailKey := config.ImportSetting()

	CONFIG_SMTP_HOST := emailKey.Host
	CONFIG_SMTP_PORT := emailKey.Port
	CONFIG_SENDER_NAME := emailKey.Name
	CONFIG_AUTH_EMAIL := emailKey.Email
	CONFIG_AUTH_PASSWORD := emailKey.Password

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "viruzboyz11@gmail.com")
	mailer.SetAddressHeader("Cc", "viruzboyz11@gmail.com", "hafiz2")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
	mailer.Attach("./sample.png")

	log.Println(CONFIG_AUTH_EMAIL)
	log.Println(CONFIG_AUTH_PASSWORD)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
