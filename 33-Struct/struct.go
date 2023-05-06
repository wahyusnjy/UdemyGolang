package main

import "fmt"



type Customer struct {
	Nama,Address string
	Age 	      int
}

func main(){
 	var yuu Customer
	yuu.Nama = "Wahyu"
	yuu.Address = "Jakarta"
	yuu.Age	= 19

	fmt.Println(yuu.Nama)
	fmt.Println(yuu.Address)
	fmt.Println(yuu.Age);

	//Struct Literal 

	joko := Customer{
		Nama: "Joko",
		Address: "Indonesia",
		Age: 30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 35}
	fmt.Println(budi)
}