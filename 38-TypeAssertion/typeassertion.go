package main 

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()

	// var resultString = result.(string)
	// fmt.Println(resultString)

	// var resultInt = result.(int)
	// fmt.Println(resultInt)
	// Result panic 

	switch value := result.(type){
	case string :
		fmt.Println("String", value)
	case int : 
		fmt.Println("Int", value)
	default : 
		fmt.Println("Unknown", value)
	}
}