package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func handleRequests() {

	//Initializing mux server to handle the request from the backend
	myRouter := mux.NewRouter().StrictSlash(true)


	//Redirect the requests coming from the frontend
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/callSession", manageSession)
	myRouter.HandleFunc("/deleteUser", deleteUser)
	myRouter.HandleFunc("/updateUser", updateUser)
	myRouter.HandleFunc("/customer/call", customerCall).Methods("POST")
	myRouter.HandleFunc("/customer/registration", customerRegistration).Methods("POST")
	myRouter.HandleFunc("/serviceprovider/registration", serviceProviderRegistration).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}