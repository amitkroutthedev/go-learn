package main

import "fmt"

func main() {
	fmt.Println("Varibales")
	var username string = "tester"
	fmt.Println("Username value:",username)
	fmt.Printf("Username type %T\n",username)

	var isLogged bool = false
	fmt.Println("isLogged value:",isLogged)
	fmt.Printf("isLogged type %T\n",isLogged)

	var smallVal int = 256
	fmt.Println("smallVal value:",smallVal)
	fmt.Printf("smallVal type %T\n",smallVal)

	var smallFloat float64 = 256.23413242341234
	fmt.Println("smallFloat value:",smallFloat)
	fmt.Printf("smallFloat type %T\n",smallFloat)

	var anotherVariable int
	fmt.Println("anotherVariable value:",anotherVariable)
	fmt.Printf("anotherVariable type %T\n",anotherVariable)


	numberOfUser:=3000
	fmt.Println(numberOfUser)
}