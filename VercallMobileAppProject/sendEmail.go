package main

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"log"
	"net/smtp"
)

func generateEmailVerificationLinkForCustomer(ctx context.Context, client *auth.Client, customDetails CustomerRegistration) {
	actionCodeSettings := &auth.ActionCodeSettings{
		URL:                   "http://localhost/",
		HandleCodeInApp:       true,
	}
	// [START email_verification_link]
	email := customDetails.Email
	link, err := client.EmailVerificationLinkWithSettings(ctx, email, actionCodeSettings)
	if err != nil {
		log.Fatalf("error generating email link: %v\n", err)
	}

	var body string = "Hello" + " "+ customDetails.Firstname+ ","+ "\n" +

		"\n" + link + "\n" + "Follow this link to verify your email address. \n" +

		"\n" + "If you didn’t ask to verify this address, you can ignore this email. \n" +

		"\n" + "Thanks, \n" +

		"\n" + "Your Vercall Mobile App Team"
	//sendCustomEmail(email, displayName, link)
	send(body , customDetails.Email)
	// [END email_verification_link]
}

func generateEmailVerificationLinkForServiceProvider(ctx context.Context, client *auth.Client, serviceProviderDetails ServiceProviderRegistration) {
	actionCodeSettings := &auth.ActionCodeSettings{
		URL:                   "http://localhost/",
		HandleCodeInApp:       true,
	}
	// [START email_verification_link]
	email := serviceProviderDetails.Email
	link, err := client.EmailVerificationLinkWithSettings(ctx, email, actionCodeSettings)
	if err != nil {
		log.Fatalf("error generating email link: %v\n", err)
	}

	// Construct email verification template, embed the link and send
	// using custom SMTP server.
	//fmt.Println("Hello + customDetails.Firstname,+
	//
	//	link + "Follow this link to verify your email address.+"
	//
	//
	//    "If you didn’t ask to verify this address, you can ignore this email."
	//
	//	"Thanks,"
	//
	//	"Your %APP_NAME% team");

	var body string = "Hello" + " "+ serviceProviderDetails.Firstname+ ","+ "\n" +

		"\n" + link + "\n" + "Follow this link to verify your email address. \n" +

		"\n" + "If you didn’t ask to verify this address, you can ignore this email. \n" +

		"\n" + "Thanks, \n" +

		"\n" + "Your Vercall Mobile App Team"
	//sendCustomEmail(email, displayName, link)
	send(body , serviceProviderDetails.Email)
	// [END email_verification_link]
}

func send(body string, To string) {
	from := "username"
	pass := "password"
	to := To

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Email Verification - Vercall Mobile Application\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, Email has been sent to the user.")
}
