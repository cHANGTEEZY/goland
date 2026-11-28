package main

// need to do something
func (e email) cost() int {
	return 1
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
