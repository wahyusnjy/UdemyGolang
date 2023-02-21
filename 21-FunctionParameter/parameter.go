package main

import (
	"fmt"
)

func SayParameter(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

func main(){
	firstName := "Wahyu"
	SayParameter(firstName, "Sanjaya")
	SayParameter("Rui","Hermawang")
}