package main

import (
	"fmt"
)

func main(){
	counter := 1
	for counter <= 10 {
		fmt.Println("Looping Ke", counter)
		counter++
	}

	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan Ke",counter2)
	}

	slice := []string{"Wahyu", "Sanjaya","Abdul","Fattah","Maulana","Fajar"}

	for i := 0; i <  len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"Wahyu", "Abdul","Maul","Udin","Rui","Begug"}
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
		// fmt.Println(name)
	} 

	person := make(map[string]string)
	person["name"] = "Wahyu"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}