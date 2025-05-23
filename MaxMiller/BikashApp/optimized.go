package main

import "fmt"

var accountBalance float64 =1000.00
func main_(){


	// var ussdInput string
	// fmt.Print("Please dial *247# to start your transaction: ")
	// fmt.Scan(&ussdInput)

	// if ussdInput != "*247#" {
	// 	fmt.Println("Invalid USSD code. Exiting...")
	// 	return
	// }

	fmt.Println("Welcome to Fake-Bikash App")
	fmt.Println("-------------------------")

	for {
		showMenu()
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			checkBalance()
		case 2:
			cashIn()
		case 3:
			cashOut()
		case 4:
			fmt.Println("-----Thanks for choosing Fake-Bikash App----")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}


func showMenu(){
	fmt.Println("What do you want to do?")
		fmt.Println("1. Check Your Balance")
		fmt.Println("2. Cash In")
		fmt.Println("3. Cash Out")
		fmt.Println("4. Exit Fake-Bikash App")
}

func checkBalance(){
	fmt.Printf("Your Balance is: %.2f Tk\n", accountBalance)
}

func cashIn() {
	var amount float64
	fmt.Print("Enter your Cash In amount: ")
	fmt.Scan(&amount)

	if !isValidAmount(amount) {
		return
	}

	accountBalance += amount
	fmt.Printf("Balance Updated! Current balance is: %.2f Tk\n", accountBalance)
}


func cashOut() {
	var amount float64
	fmt.Print("Enter your Cash Out amount: ")
	fmt.Scan(&amount)

	if !isValidAmount(amount) {
		return
	}

	if amount > accountBalance {
		fmt.Println("Insufficient balance!")
		return
	}

	accountBalance -= amount
	fmt.Printf("Balance Updated! Current balance is: %.2f Tk\n", accountBalance)
}

func isValidAmount(amount float64) bool {
	if amount <= 0 {
		fmt.Println("Invalid amount! Please enter an amount greater than 0.")
		return false
	}
	return true
}