package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer")
	//var ptr *int
	//fmt.Println("Value of default pointer",ptr)
	//default value nil
	myNumber:=23

	var ptrMyNumber = &myNumber
	fmt.Println("memory Value of pointer",ptrMyNumber)
	fmt.Println("Actual value of pointer",*ptrMyNumber)

	*ptrMyNumber = *ptrMyNumber * 2
	fmt.Println("New value is",*ptrMyNumber)
}
