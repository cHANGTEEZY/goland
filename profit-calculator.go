package main

import (
	"errors"
	"fmt"
	"os"
)

func storeResultsInFile(fileName string, value float64) error {
	data := fmt.Sprintf("%.3f\n", value)
	return os.WriteFile(fileName+".txt", []byte(data), 0644)
}

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)


	_ = storeResultsInFile("profit", profit)
	_ = storeResultsInFile("ebt", ebt)
	_ = storeResultsInFile("ratio", ratio)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)

	var ratio float64
	if profit != 0 {
		ratio = ebt / profit
	}

	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return 0, errors.New("invalid input type, please enter a number")
	}

	if userInput <= 0 {
		return 0, errors.New("please enter a number greater than 0")
	}

	return userInput, nil
}
