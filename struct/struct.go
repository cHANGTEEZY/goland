package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	userFirstName string
	userLastName  string
	userBirthDate string
	userCreatedAt time.Time
}

// * constructor function for creating a new user, using pointers to avoid copying the entire struct
// * and return a pointer to the new user
func createUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, Last name and Birthdate is required")
	}

	return &user{
		userFirstName: firstName,
		userLastName:  lastName,
		userBirthDate: birthdate,
		userCreatedAt: time.Now(),
	}, nil
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

	var appUser *user

	appUser, err := createUser(firstName, lastName, birthdate)

	if (err != nil) {
		fmt.Println(err)
		return
	}

	// ? ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserDetails()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
