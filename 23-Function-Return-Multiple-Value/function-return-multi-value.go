package main

import "fmt"

func getFullName() (string,string,string,string) {
	return "Wahyu", "Sanjaya" , "Ira", "Ramadhanty"
}

func main(){
	firstName, _ , ayangFirstName , ayangLastName  := getFullName()
	fmt.Println(firstName)
	fmt.Println(ayangFirstName, ayangLastName)
}