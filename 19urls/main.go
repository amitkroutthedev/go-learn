package main

import (
	"fmt"
	"net/url"
)

const URL string="https://www.akrout.dev:3000/learn?coursename=reactjs&paymentid=321ssw221" 

func main() {
	fmt.Println("Welcome to handling URLS")
    fmt.Println("URL:",URL)

	//parsing
	result,_:=url.Parse(URL)

	fmt.Println("Parsing Result Scheme:",result.Scheme)
	fmt.Println("Parsing Result Host:",result.Host)
	fmt.Println("Parsing Result Path:",result.Path)
	fmt.Println("Parsing Result Port:",result.Port())
	fmt.Println("Parsing Result RawQuery:",result.RawQuery)

	qparams:=result.Query()
	fmt.Printf("Type of Query params %T",qparams)
	fmt.Println("\nQuery",qparams)
 
	for _,val:=range qparams{
		fmt.Println("Param is",val)
	}
    
	partsOfUrl:=&url.URL{
		Scheme: "https",
		Host:"akrout.dev",
		Path:"/tutcss",
		RawPath: "url=hitesh",
	}

	anotherURL:=partsOfUrl.String()

	fmt.Println(anotherURL)
}
