package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage of Vercall Mobile Application!")
	fmt.Println("Endpoint Hit: homePage")
}