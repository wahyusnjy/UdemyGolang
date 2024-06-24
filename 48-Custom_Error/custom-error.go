package main

import (
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "wahyu" {
		return &notFoundError{"data not found"}
	}

	//data Ditemukan

	return nil
}


func main() {
	err := SaveData("wahyu", nil)
	if err != nil {
		//terjadi error
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error: ", validationErr.Error())
		// }else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error : ", notFoundErr.Error())
		// }else {
		// 	fmt.Println("Unknown Error: ", err.Error())
		// }

		//terjadi error dengan switch
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation error: ", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error : ", finalError.Error())
		default:
			fmt.Println("unknown error : ", finalError.Error())
		}
	}else {
		//success
		fmt.Println("Sukses")
	}
}