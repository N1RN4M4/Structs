package main

import "fmt"

func main() {
	firstName := getUserInfo("First Name: ")
	lastName := getUserInfo("Last Name: ")
	birthDate := getUserInfo("Birth Date: ")

	fmt.Println(firstName, lastName, birthDate)
}

// Need to pass te data
func OutputUserInfo(firstName, lastName, birthDate string) {
	fmt.Println(firstName, lastName, birthDate)
}

func getUserInfo(userInput string) string {
	fmt.Print(userInput)
	var value string
	fmt.Scan(&value)
	return value
}
