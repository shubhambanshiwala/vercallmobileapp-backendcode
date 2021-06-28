package main

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"net/http"
	"context"
)

func updateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Update User API is in execution")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var updateUser UpdateUser;
	json.Unmarshal(reqBody, &updateUser)
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
	iter := dbClient.Collection("customer_registration").Where("email", "==",  updateUser.Email).Documents(context.Background())

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

			_, err  = dbClient.Collection("customer_registration").Doc(doc.Ref.ID).Set(context.Background(), map[string]interface{}{
				"city": updateUser.City,
				"address": updateUser.Address,
				"date_of_birth": updateUser.DateOfBirth,
				"email": updateUser.Email,
				"firstName": updateUser.Firstname,
				"lastName": updateUser.Lastname,
				"phone_number": updateUser.PhoneNumber,
				"post_code": updateUser.PostCode,
			}, firestore.MergeAll)
			if err != nil {
				// Handle any errors in an appropriate way, such as returning them.
				log.Printf("An error has occurred: %s", err)
			}
			// handle err..
		}
	}
}
