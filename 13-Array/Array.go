package main

import "fmt"

func main() {
	//Array manual
	var names [3]string

	names[0] = "Wahyu"
	names[1] = "Sanjaya"
	names[2] = "Ramadhan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//aray praktis
	var values = [3]int{
		80,
		85,
		90,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))
}
