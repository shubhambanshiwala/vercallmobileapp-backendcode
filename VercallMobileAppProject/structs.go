package main

import "encoding/xml"

type CustomerRegistration struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	RegistrationStatus string `json:"registrationstatus"`
	City string `json:"city"`
	DateOfBirth string `json:"dataOfBirth"` //--date of birth
	PostCode int `json:"postCode"`
	PhoneNumber string `json:"phoneNumber"`
}

type UpdateUser struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	RegistrationStatus string `json:"registrationstatus"`
	City string `json:"city"`
	DateOfBirth string `json:"dataOfBirth"` //--date of birth
	PostCode int `json:"post_code"`
	PhoneNumber string `json:"phone_number"`
}

type CallSession struct {
	Email string `json: "email"`
	Status string `json: "status"`
}

type DeleteUser struct {
	Email string `json: "email"`
}

type Article struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type ServiceProviderRegistration struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	City string `json:"city"`
	DateOfBirth string `json:"dataOfBirth"` //--date of birth
	PostCode string `json:"post_code"`
	PhoneNumber string `json:"phoneNumber"`
	Company_name string `json:"companyName"`
	Abn string `json:"abn"`
}

type CustomerLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Status struct{
	Status string `json:"status"`
	MessageCode int `json:"messageCode"`
}
type ServiceProviderLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
type TwiML struct {
	XMLName xml.Name `xml:"Response"`
	Say    string `xml:",omitempty"`
	Dial   Dial
}

type Dial struct {
	Number string
}