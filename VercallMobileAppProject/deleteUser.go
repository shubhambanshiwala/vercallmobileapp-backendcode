package main

import (
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

func deleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Delete User API is in execution")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var deleteUser DeleteUser;
	json.Unmarshal(reqBody, &deleteUser)
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
	iter := dbClient.Collection("customer_registration").Where("email", "==",  deleteUser.Email).Documents(context.Background())

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

			_, err  = dbClient.Collection("customer_registration").Doc(doc.Ref.ID).Delete(context.Background())
			if err != nil {
				// Handle any errors in an appropriate way, such as returning them.
				log.Printf("An error has occurred: %s", err)
			}
			// handle err..
		}
	}
}
