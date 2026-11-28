package main

// setting cost of sending email based on wheather user is subscribed or not
func (e email) cost() int {
	costPerCharacter := 2

	if !e.isSubscribed {
		costPerCharacter = 5
		return costPerCharacter
	}
	return costPerCharacter
}

// need to do something
func (e email) format() string {
	return ""
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
