package main

import (
	"fmt"
	"os"
)

func storeResultsInFile(fileName string, value float64) error {
	data := fmt.Sprintf("%.3f\n", value)
	return os.WriteFile(fileName+".txt", []byte(data), 0644)
}

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	storeResultsInFile("profit",profit )
	storeResultsInFile("ebt",ebt )
	storeResultsInFile("ratio",ratio )

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	_,err := fmt.Scan(&userInput)

	if err != nil {
		panic("invalid input, please enter a valid input")
	}

	return userInput
}
