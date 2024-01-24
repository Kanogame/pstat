package utils

import "fmt"

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleClientError(err error, message string) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
