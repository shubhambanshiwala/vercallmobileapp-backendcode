package main

import (
	"cloud.google.com/go/firestore"
	"time"

	//"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"net/http"
)


func manageSession(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ManageSession API is in execution");
	reqBody, _ := ioutil.ReadAll(r.Body)
	var callSession CallSession;
	json.Unmarshal(reqBody, &callSession)
	opt := option.WithCredentialsFile("C:/Users/Shubham/Desktop/Uni-Project/Security Key/vercall-mobile-app-firebase-adminsdk-z7jzd-2b93588160.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	dbClient, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}
	//var databaseData *firestore.DocumentSnapshot
	iter := dbClient.Collection("customer_registration").Where("email", "==",  callSession.Email).Documents(context.Background())

	if (iter == nil){
		fmt.Println("There is no data with email id")
	} else  {
		fmt.Println("There is data exist with email id")
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("error creating user: %v\n", err)
			}

			_, err  = dbClient.Collection("customer_registration").Doc(doc.Ref.ID).Update(context.Background(), []firestore.Update{
				{
					Path:  "registrationstatus",
					Value: "true",
				},
			})
			// handle err..

		}
	}
	// Calling NewTimer method
	newtimer := time.NewTimer(60* time.Second)
	<-newtimer.C
	dbdata := dbClient.Collection("customer_registration").Where("email", "==",  callSession.Email).Documents(context.Background())
	for {
		doc, err := dbdata.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error creating user: %v\n", err)
		}
		_, err  = dbClient.Collection("customer_registration").Doc(doc.Ref.ID).Update(context.Background(), []firestore.Update{
			{
				Path:  "registrationstatus",
				Value: "false",
			},
		})
	}
}