package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var noktpwahyu NoKtp = "3140129319"
	var status Married = true
	
	fmt.Println(noktpwahyu)
	fmt.Println(status)
}