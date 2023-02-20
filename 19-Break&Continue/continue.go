package main 

import (
	"fmt"
)

func main(){
	for c := 0; c < 10; c++ {
		if c % 2 == 0 {
			continue
		}
		fmt.Println("Perulangan Ke", c)
	}
}