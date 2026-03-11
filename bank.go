package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
		//panic("Cant continue.")
	}

	fmt.Println("Welcome to Bryn Bank")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		//fmt.Println("Your choice is:", choice)
		switch choice {
		case 1:
			fmt.Println("Your Balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("How much you want to deposit? ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				println("Invalid Amount, must be greater then 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("You balance now is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdrawAmount float64
			fmt.Print("How much you want to withdraw? ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount, must be greater then 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount, You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("You balance now is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
