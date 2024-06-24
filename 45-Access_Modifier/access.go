package main

import (
	"udemyGolang/45-Access_Modifier/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Wahyu")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses
	// fmt.Println(helper.sayGoodbye("wahyu")) // tidak bisa di akses

}