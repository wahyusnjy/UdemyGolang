package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye" + name
}

func main(){
	goodbye := getGoodBye

	result := goodbye(" Wahyu")

	fmt.Println(result)
	fmt.Println(getGoodBye(" Yuu"))
}