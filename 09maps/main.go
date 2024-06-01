package main

import "fmt"

func main() {
	fmt.Println("Maps in go lang")

	languages:= make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages,"RB")
	fmt.Println(languages)
   
	for key,value:=range languages{
		fmt.Printf("For key %v,value is %v\n",key,value)
	}
	for _,value:=range languages{
		fmt.Printf("value is %v\n",value)
	}

}
