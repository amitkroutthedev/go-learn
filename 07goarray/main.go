package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Tomato"

	fmt.Println("Fruit list is",fruitList)
	fmt.Println("Length Fruit list is",len(fruitList))

	var vegList = [3]string{"potato","beans","mushroom"}
	fmt.Println("Vegies list is",vegList,len(vegList))
}
