package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string{
	return animal.Name
}

func main(){
	var wahyu Person
	wahyu.Name = "Wahyu"

	SayHello(wahyu)

	cat := Animal{
		Name: "Push",
	}

	SayHello(cat)
}