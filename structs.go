package main

import "fmt"

func main() {
	// firstName //will take from the new func
	firstName := getUserInfo("First Name: ")
	// lastName
	lastName := getUserInfo("Last Name: ")
	// birthDate
	birthDate := getUserInfo("Birth Date: ")

	fmt.Println(firstName, lastName, birthDate)
}

func getUserInfo(userInput string) string {
	fmt.Print(userInput)
	var value string
	fmt.Scan(&value)
	return value
}
