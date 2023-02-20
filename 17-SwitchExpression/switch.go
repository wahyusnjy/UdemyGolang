package main

import "fmt"

func main(){
	name := "Mansur"

//1
	switch name {
	case "Wahyu":
		fmt.Println("Hello Wahyu")
		fmt.Println("Case 1 Success")
	case "Sanjaya":
		fmt.Println("Hello Sanjaya")
		fmt.Println("Case 2 Success")
	default:
		fmt.Println("Boleh Tau siapa namanya?")
		fmt.Println("Case 3 Success")
	}
 //2
	// switch length := len(name); length > 5{
	// case true :
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false :
	// 	fmt.Println("Nama Sudah Benar")
	// }

//3 
	panjang := len(name)
	switch  {
	case panjang > 10:
		fmt.Println("Nama Terlalu Panjang")
	case panjang > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}