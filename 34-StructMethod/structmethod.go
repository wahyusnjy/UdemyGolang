package main

import "fmt"

type SayHi struct {
	Nama,Address string
	Age 	      int
}

func (sayhi SayHi) sayHello(){
	fmt.Println("Hello My Name is", sayhi.Nama)
}

func (a SayHi) sayHuu(){
	fmt.Println("Huuuu from", a.Nama)
}

func main(){
	rully := SayHi{Nama: "Rully"}
	rully.sayHello()

	wahyu := SayHi{Nama: "Wahyu"}
	wahyu.sayHello()
	wahyu.sayHuu()
}