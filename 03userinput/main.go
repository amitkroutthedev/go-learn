package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welome :="welcome to user input"
	fmt.Println(welome)

	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter that rating:")

	//comma ok syntax || error ok syntac

	input,_:=reader.ReadString('\n')

	fmt.Println("Rating ",input)

}
