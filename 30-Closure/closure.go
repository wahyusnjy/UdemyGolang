package main

import "fmt"

func main() {
	counter := 0
	name := "Wahyu"

	increment := func ()  {
		name = "Rara"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
} 