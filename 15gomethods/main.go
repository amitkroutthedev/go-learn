package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")
	// no inheritance in golang; No super or parent
	amit := User{"Amit", "akrout455@gmail.com", true, 24}
	fmt.Println("User details", amit)
	fmt.Printf("User: %+v\n", amit)
	fmt.Printf("Name is %v and email is %v\n", amit.Name, amit.Email)
	amit.GetStatus()
	amit.NewMail()
	
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	//oneAge int   //not exportable not caps at first
}	

func (u User) GetStatus(){
fmt.Println("Status",u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Mail",u.Email)
}