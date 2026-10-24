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

// constructor function for creating a new user
func createUser(firstName, lastName, birthdate string) user {
	return user{
		userFirstName: firstName,
		userLastName:  lastName,
		userBirthDate: birthdate,
		userCreatedAt: time.Now(),
	}
}

// * adding a method to the user struct, and the appuser here is called the receiver
func (appUser user) outputUserDetails() {
	fmt.Println(appUser.userFirstName, appUser.userLastName, appUser.userBirthDate)

	fmt.Println("User created at:", appUser.userCreatedAt)
}

func (appUser *user) clearUserDetails() {
	appUser.userFirstName = ""
	appUser.userLastName = ""
	appUser.userBirthDate = ""
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = createUser(firstName, lastName, birthdate)

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserDetails()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
