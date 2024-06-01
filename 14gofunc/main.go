package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to functions")

	result:=additionres(4,5)

	fmt.Println("Result is",result)

	proRes:=proAdder(2,5,6,7,1,3,5)
	fmt.Println(proRes)
}

//proRes,myMessage:=proAdder(2,5,6,7,1,3,5)

/*
func proAdder(values ...int) (int,string){
	total:=0
	for _,val:=range values{
		total+=val
	}
	return total,"Extra string"
}
*/

func proAdder(values ...int) int{
	total:=0
	for _,val:=range values{
		total+=val
	}
	return total
}

func additionres(valOne int,valTwo int) int{
	return valOne+valTwo
}


func greeter(){
	fmt.Println("Hello greeter")
}