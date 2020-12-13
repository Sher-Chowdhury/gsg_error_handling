package main

import "fmt"

func lessThanTen(i int) (response string, error string) {

	if i < 10 {

		response = "The number is less than 10"
		error = ""
	} else {

		response = "The number is too big"
		error = "ERROR: hit an error"

	}
	return response, error

}

func main() {

	result, err := lessThanTen(5)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}