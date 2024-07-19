package main

import (
	"flag"
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"net/smtp"
	"os"
)

func main() {
	useEmail := flag.Bool("email", false, "notify via email")
	useSMS := flag.Bool("sms", false, "notify via SMS")
	if *useEmail {
		sendEmail()
	}
	if *useSMS {
		sendSMS()
	}
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

func sendSMS() {
	client := twilio.NewRestClient()
	toPhoneNumber, ok := os.LookupEnv("TO_PHONE_NUMBER")
	if !ok {
		log.Println("TO_PHONE_NUMBER environment variable not set")
		return
	}
	twilioPhoneNumber, ok := os.LookupEnv("TWILIO_PHONE_NUMBER")
	if !ok {
		log.Println("TWILIO_PHONE_NUMBER environment variable not set")
		return
	}

	params := &openapi.CreateMessageParams{}
	params.SetTo(toPhoneNumber)
	params.SetFrom(twilioPhoneNumber)
	params.SetBody("Hello from Golang!")

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}

}
