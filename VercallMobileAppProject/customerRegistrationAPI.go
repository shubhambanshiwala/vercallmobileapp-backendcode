package main

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func customerRegistration(w http.ResponseWriter, r *http.Request){

	//location, err := time.LoadLocation("Australia/Melbourne")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customDetails CustomerRegistration
	json.Unmarshal(reqBody, &customDetails)

	opt := option.WithCredentialsFile("C:/Users/Shubham/Desktop/Uni-Project/Security Key/vercall-mobile-app-firebase-adminsdk-z7jzd-2b93588160.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	ctx := context.Background()
	authClient, err := app.Auth(context.Background())
	fmt.Println(customDetails.PhoneNumber);
	params := (&auth.UserToCreate{}).
		Email(customDetails.Email).
		EmailVerified(false).
		PhoneNumber("+61484021509").
		Password(customDetails.Password).
		DisplayName(customDetails.Firstname)
	u, err := authClient.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)
	generateEmailVerificationLinkForCustomer(ctx, authClient, customDetails)
	send("hello there", "string")
	dbClient, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}
	layout := "2006-01-02 15:04:05"
	currentTime, _ := time.Parse(layout, time.Now().Format("2006-01-02 15:04:05"));
	dateOfBirth, _ := time.Parse("02-01-2006", customDetails.DateOfBirth)
fmt.Println(customDetails.Email)
	_, _, err = dbClient.Collection("customer_registration").Add(context.Background(), map[string]interface{}{
		"firstName": customDetails.Firstname ,
		"lastName":  customDetails.Lastname,
		"email": customDetails.Email,
		"address":  customDetails.Address,
		"timestamp_registered": currentTime.UTC(),
		"city": customDetails.City,
		"date_of_birth": dateOfBirth,
		"post_code": customDetails.PostCode,
		"phone_number": "+61484021509",
		"registrationstatus": "false",
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

}


