package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	fmt.Println("Vercall APP Backend API Running.....")

	//Method to route the inbound request
	handleRequests()
}