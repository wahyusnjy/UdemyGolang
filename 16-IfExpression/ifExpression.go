package main

import "fmt"

// "fmt"
func main (){
	var name = "Fredrin Sambo"

	if name == "wahyu" {
		fmt.Println("Hello Wahyu")
	}  else if name == "sanjaya" {
		fmt.Println("Hello Wahyu Sanjaya")
	}	else if name == "Fredrin Sambo" {
		fmt.Println("Pidana MATI :D")
	}  else {
		fmt.Println("Lu siapa ? ")
	}

	if length := len(name);length > 5 {
		fmt.Println("Terlalu Edun")
	} else {
		fmt.Println("Sudah Benar")
	}

}