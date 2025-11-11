package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserInfo("First Name: ")
	userLastName := getUserInfo("Last Name: ")
	userBirthDate := getUserInfo("Birth Date: ")

	fmt.Println(firstName, lastName, birthDate)

	var appUser User

	appUser = User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}

	OutputUserInfo(appUser)

}

// Need to pass the data
func OutputUserInfo(u User) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserInfo(userInput string) string {
	fmt.Print(userInput)
	var value string
	fmt.Scan(&value)
	return value
}
