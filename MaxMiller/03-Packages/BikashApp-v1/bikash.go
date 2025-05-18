package main

import (
	"fmt"
)



func main() {

	accountBalance := 1000.00;


	fmt.Println("Welcome to Fake-Bikash App")
	fmt.Println("-------------------------")

	for {

		PresentOptions()

		var choice int
		fmt.Print("Enter your option: ")
		fmt.Scan(&choice)
		// fmt.Println("You have selected option:", choice)

		switch choice {
		case 1:
			fmt.Printf("You Balance is : %.2f Tk\n", accountBalance)
		case 2:
			//user cash in
			fmt.Print("Enter your Cash in amount: ")
			var cashInAmount float64
			fmt.Scan(&cashInAmount)
			fmt.Printf("You entered: %.2f Tk\n", cashInAmount)

			if cashInAmount <= 0 {
				fmt.Println("Invalid amount! Please enter a amount greater than 0.")
				// return;
				continue
			}

			accountBalance = accountBalance + cashInAmount
			fmt.Printf("Balanced Updated!Current balance is: %.2f Tk\n", accountBalance)
		case 3:
			//user cash out
			fmt.Print(("Enter your Cash out amount: "))
			var cashOutAmount float64
			fmt.Scan((&cashOutAmount))

			if cashOutAmount <= 0 {
				fmt.Println("Invalid amount! Please enter a amount greater than 0.")
				continue
			}

			if cashOutAmount > accountBalance {
				fmt.Println("Insufficient balance!")
				// return
				continue
			}

			accountBalance = accountBalance - cashOutAmount
			fmt.Printf("Balanced Updated!Current balance is: %.2f Tk\n", accountBalance)
		default:
			fmt.Println("-----Thanks for chosing Fake-Bikash App----")
			return

		}

	}

}