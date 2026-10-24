package main

import (
	"fmt"
	"go-foundation/fileops/getFloatFromFile"
	"os"
)

const balanceFile = "balance.txt"
const withdrawReceiptFile = "withdrawReceipt.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func writeWithdrawReceiptToFile(amount float64) {
	receiptText := fmt.Sprintf("%.2f", amount)
	os.WriteFile(withdrawReceiptFile, []byte(receiptText), 0644)
}

func main() {
	fmt.Println("Welcome to GO Bank!")

	for {
		accountBalance, err := getFloatFromFile.GetFloatFromFile()
		if err != nil {
			fmt.Println("Error loading balance:", err)
			panic(err)
		}

		var choice int

		presentOptions()

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your current balance is: $%.2f\n", accountBalance)

		case 2:
			fmt.Print("How much do you want to deposit? ")
			var depositAmount float64
			fmt.Scanln(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Deposit successful! New balance: $%.2f\n", accountBalance)

		case 3:
			fmt.Print("How much do you want to withdraw? ")
			var withdrawAmount float64
			fmt.Scanln(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds!")
				continue
			}

			accountBalance -= withdrawAmount
			writeBalanceToFile(accountBalance)
			writeWithdrawReceiptToFile(withdrawAmount)
			fmt.Printf("Withdrawal successful! New balance: $%.2f\n", accountBalance)

		case 4:
			fmt.Println("Thank you for using GO Bank! Goodbye!")
			return

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}
}
