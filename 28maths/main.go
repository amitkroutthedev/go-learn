package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")
	//p
	//rand.Seed(time.Now().UnixNano())
	//x:=rand.Intn(10)
	//fmt.Println("math rand number",x)
	//random from  crypto
	myRandomNumber,_:=rand.Int(rand.Reader,big.NewInt(20))
	fmt.Println("rand number",myRandomNumber)
}
