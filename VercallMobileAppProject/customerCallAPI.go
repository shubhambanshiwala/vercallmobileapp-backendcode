package main

import (
	"context"
	"encoding/xml"

	//"encoding/xml"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"net/http"
)
var c CustomerRegistration
func customerCall(w http.ResponseWriter, r *http.Request)  {

	w.Write([]byte("Welcome to Vercall Mobile Communication Services.  Please wait! while we are authenticating you"))
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println("Caller Number", r.Form.Get("Caller"))
	log.Println("Caller SID - Uniquet ID", r.Form.Get("CallSid"))
	log.Println("AccountSid", r.Form.Get("AccountSid"))
	opt := option.WithCredentialsFile("C:/Users/Shubham/Desktop/Uni-Project/Security Key/vercall-mobile-app-firebase-adminsdk-z7jzd-2b93588160.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	dbClient, err := app.Firestore(context.Background())
	//var databaseData *firestore.DocumentSnapshot
	iter := dbClient.Collection("customer_registration").Where("phone_number", "==", r.Form.Get("Caller") ).Documents(context.Background())
	for {
		databaseData, err := iter.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {

		}
		fmt.Println(databaseData.Data())
		fmt.Println(databaseData.DataTo(&c))
		fmt.Println(c.RegistrationStatus)
	}
			if (c.RegistrationStatus == "false") {
			twiml := TwiML{Say: "Authentication is failed, please call by logging to your vercall app again."}
			x, err := xml.MarshalIndent(twiml, "", " ")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/xml")
			fmt.Println(x)
			w.Write(x)
		} else if (c.RegistrationStatus == "true") {
			dial := Dial{Number: "+61-4-111-40-360"}
			twiml := TwiML{Say: "Authentication is successfull, your call will be forwarded to", Dial: dial}
			x, err := xml.MarshalIndent(twiml, "", " ")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/xml")
			fmt.Println(x)
			w.Write(x)
		} else {
			twiml := TwiML{Say: "Your authentication cannot be proceeded since there is no data available, please contact admin"}
			x, err := xml.MarshalIndent(twiml, "", " ")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/xml")
			fmt.Println(x)
			w.Write(x)
		}

}