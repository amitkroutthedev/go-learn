package main

import "fmt"

func main() {
	defer fmt.Println("Hello defer")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Welcome to defer")
	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}