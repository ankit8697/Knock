package main

import (
	"log"
	"net/smtp"
	"os"
)

func main() {
	sendEmail()
}

// sendEmail sends an email to the address in the environment variable. If either KNOCK_EMAIL_ADDRESS or
// KNOCK_EMAIL_APP_PASSWORD are not present, we will not send the email.
func sendEmail() {
	email, ok := os.LookupEnv("KNOCK_EMAIL_ADDRESS")
	if !ok {
		log.Println("If you would like to get notified via email, please set the KNOCK_EMAIL_ADDRESS environment" +
			" variable to your email address.")
		return
	}
	pass, ok := os.LookupEnv("KNOCK_EMAIL_APP_PASSWORD")
	if !ok {
		log.Println("If you would like to get notified via email, please set the KNOCK_EMAIL_APP_PASSWORD " +
			"environment variable to the app password of your email address.")
		return
	}

	msg := "From: " + email + "\n" +
		"To: " + email + "\n" +
		"Subject: Your command is done executing!\n\n" +
		"Time email go see what happened email your command!"

	auth := smtp.PlainAuth("", email, pass, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587",
		auth,
		email, []string{email}, []byte(msg))

	if err != nil {
		log.Printf("Something went wrong when sending the email: %s", err)
		return
	}
	log.Println("Successfully sent email " + email)
}
