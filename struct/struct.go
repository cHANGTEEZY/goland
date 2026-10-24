package main

import (
	"fmt"
	"time"
)

type user struct {
	userFirstName string
	userLastName  string
	userBirthDate string
	userCreatedAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		userFirstName: firstName,
		userLastName:  lastName,
		userBirthDate: birthdate,
		userCreatedAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	outputUserDetails(&appUser)
}

func outputUserDetails(appUser *user) {
	fmt.Println(appUser.userFirstName, appUser.userLastName, appUser.userBirthDate)

	fmt.Println("User created at:", appUser.userCreatedAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
