package main

import "fmt"

func (e email) cost() int {
	costPerCharacter := 2
	message := e.body

	if !e.isSubscribed {
		costPerCharacter = 5
		return len(message) * costPerCharacter
	}

	return len(message) * costPerCharacter
}

func (e email) format() string {
	message := "Not Subscribed"
	if e.isSubscribed {
		message = "Subscribed"
	}

	return fmt.Sprintf("'%s' | %s", e.body, message)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
