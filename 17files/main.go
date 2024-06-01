package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file."

	file, err := os.Create("./testfile.txt")
	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	/*if err != nil {
		panic(err)
	}*/
	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./testfile.txt")
}

func readFile(filename string) {
	databytes, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is \n", databytes, "\n", string(databytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
