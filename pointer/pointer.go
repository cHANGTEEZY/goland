package main

import "fmt"

func pringAdultAge(age int) {
	age := 30

	fmt.Println("Age:", age)

	adultYears := getAdultAge(age)

	fmt.Println("Adult years:", adultYears)
}

func getAdultAge(age int) int {
	return age - 18
}
