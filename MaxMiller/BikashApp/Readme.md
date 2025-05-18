🔲 Run: 
- `cd BikashApp` 
- `go run bikash.go` 

I want to make a simple **bikash app** which we see in our real life
by typing `*247#`. 

<details>
  <summary> 🔲 <b>Final full code</b> </summary>

```go
// Run: go run bikash.go
package main

import "fmt"

func main() {

	accountBalance := 1000.00

	

	// Ask user to dial *247# before starting
	var ussdInput string
	fmt.Print("Please dial *247# to start your transaction: ")
	fmt.Scan(&ussdInput)

	if ussdInput != "*247#" {
		fmt.Println("Invalid USSD code. Exiting...")
		return
	}

	fmt.Println("Welcome to Fake-Bikash App")
	fmt.Println("-------------------------")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Your Balance")
		fmt.Println("2. Cash In")
		fmt.Println("3. Cash Out")
		fmt.Println("4. Exit Fake-Bikash App")

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
```
</details>

output:
```shell

❯ go run bikash.go
Please dial *247# to start your transaction: 77878
Invalid USSD code. Exiting...
❯ clear
❯ go run bikash.go
Please dial *247# to start your transaction: *247#
Welcome to Fake-Bikash App
-------------------------
What do you want to do?
1. Check Your Balance
2. Cash In
3. Cash Out
4. Exit Fake-Bikash App
Enter your option: 1
You Balance is : 1000.00 Tk
What do you want to do?
1. Check Your Balance
2. Cash In
3. Cash Out
4. Exit Fake-Bikash App
Enter your option: 2
Enter your Cash in amount: 200
You entered: 200.00 Tk
Balanced Updated!Current balance is: 1200.00 Tk
What do you want to do?
1. Check Your Balance
2. Cash In
3. Cash Out
4. Exit Fake-Bikash App
Enter your option: 1
You Balance is : 1200.00 Tk
What do you want to do?
1. Check Your Balance
2. Cash In
3. Cash Out
4. Exit Fake-Bikash App
Enter your option: 3
Enter your Cash out amount: -90
Invalid amount! Please enter a amount greater than 0.
What do you want to do?
1. Check Your Balance
2. Cash In
3. Cash Out
4. Exit Fake-Bikash App
Enter your option: 3
Enter your Cash out amount: 1100
Balanced Updated!Current balance is: 100.00 Tk
What do you want to do?
1. Check Your Balance
2. Cash In
3. Cash Out
4. Exit Fake-Bikash App
Enter your option: 1
You Balance is : 100.00 Tk
What do you want to do?
1. Check Your Balance
2. Cash In
3. Cash Out
4. Exit Fake-Bikash App
Enter your option: 4
-----Thanks for chosing Fake-Bikash App----
```


<b>Step by setp how i devlop the app.

1. Basic task
2. Next I implemnt if-else based on user choice.
3. Add Vaditation for cash in and cash out part
4. Without close the app repeated the code using for loop 
5. Add infinite time to run the loop and add condition for loop break. 
6. Next implement switch insted of nesting if else 




----


1. Basic task

```go
package main

import "fmt"

func main(){
	fmt.Println("Welcome to Bikash App");
	fmt.Println("What do you want to do?");
	fmt.Println("1. Check Your Balance");
	fmt.Println("2. Cash In");
	fmt.Println("3. Cash Out");
	fmt.Println("4. Exit Bikash App");
}
```
2. Next I implemnt if-else based on user choice.

```go
//Run: go run bikash.go

package main

import "fmt"

func main() {

	accountBalance :=1000.00;

	fmt.Println("Welcome to Fake-Bikash App")
	fmt.Println("-------------------------")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Your Balance")
	fmt.Println("2. Cash In")
	fmt.Println("3. Cash Out")
	fmt.Println("4. Exit Fake-Bikash App")

	var choice int
	fmt.Print("Enter your option: ")
	fmt.Scan(&choice)
	// fmt.Println("You have selected option:", choice)

	if choice == 1 {
		fmt.Printf("You Balance is : %.2f Tk\n", accountBalance);
	}else if choice == 2 {

		//user cash in
		fmt.Print("Enter your Cash in amount: ");
		var cashInAmount float64;
		fmt.Scan(&cashInAmount,)
		fmt.Printf("You entered: %.2f Tk\n", cashInAmount)

		accountBalance=accountBalance+cashInAmount;
		fmt.Printf("Balanced Updated!Current balance is: %.2f Tk\n", accountBalance);
	}else if choice==3 {
		fmt.Print(("Enter your Cash out amount: "))
		var cashOutAmount float64;
		fmt.Scan((&cashOutAmount));

		accountBalance=accountBalance-cashOutAmount;
		fmt.Printf("Balanced Updated!Current balance is: %.2f Tk\n", accountBalance);
	}else{
		fmt.Println("Exiting Fake-Bikash App")
		
	}
}
```
3. Add Vaditation for cash in and cash out part


```go
// Run: go run bikash.go
package main

import "fmt"

func main() {

	accountBalance := 1000.00

	fmt.Println("Welcome to Fake-Bikash App")
	fmt.Println("-------------------------")

	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Your Balance")
	fmt.Println("2. Cash In")
	fmt.Println("3. Cash Out")
	fmt.Println("4. Exit Fake-Bikash App")

	var choice int
	fmt.Print("Enter your option: ")
	fmt.Scan(&choice)
	// fmt.Println("You have selected option:", choice)

	if choice == 1 {
		fmt.Printf("You Balance is : %.2f Tk\n", accountBalance)
	} else if choice == 2 {

		//user cash in
		fmt.Print("Enter your Cash in amount: ")
		var cashInAmount float64
		fmt.Scan(&cashInAmount)
		fmt.Printf("You entered: %.2f Tk\n", cashInAmount)

		if cashInAmount <= 0 {
			fmt.Println("Invalid amount! Please enter a amount greater than 0.")
			return
		}

		accountBalance = accountBalance + cashInAmount
		fmt.Printf("Balanced Updated!Current balance is: %.2f Tk\n", accountBalance)
	} else if choice == 3 {
		fmt.Print(("Enter your Cash out amount: "))
		var cashOutAmount float64
		fmt.Scan((&cashOutAmount));


		if cashOutAmount <=0{
			fmt.Println("Invalid amount! Please enter a amount greater than 0.")
		};

		if cashOutAmount > accountBalance {
			fmt.Println("Insufficient balance!")
			return
		}

		accountBalance = accountBalance - cashOutAmount
		fmt.Printf("Balanced Updated!Current balance is: %.2f Tk\n", accountBalance)
	} else {
		fmt.Println("Exiting Fake-Bikash App")

	}
}

```


4. Add for loop so that i can selete option as much as i need

```go
// Run: go run bikash.go
package main

import "fmt"

func main() {

	accountBalance := 1000.00

	for i := 0; i < 5; i++ {
		fmt.Println("Welcome to Fake-Bikash App")
		fmt.Println("-------------------------")

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Your Balance")
		fmt.Println("2. Cash In")
		fmt.Println("3. Cash Out")
		fmt.Println("4. Exit Fake-Bikash App")

		var choice int
		fmt.Print("Enter your option: ")
		fmt.Scan(&choice)
		// fmt.Println("You have selected option:", choice)

		if choice == 1 {
			fmt.Printf("You Balance is : %.2f Tk\n", accountBalance)
		} else if choice == 2 {

			//user cash in
			fmt.Print("Enter your Cash in amount: ")
			var cashInAmount float64
			fmt.Scan(&cashInAmount)
			fmt.Printf("You entered: %.2f Tk\n", cashInAmount)

			if cashInAmount <= 0 {
				fmt.Println("Invalid amount! Please enter a amount greater than 0.")
				return
			}

			accountBalance = accountBalance + cashInAmount
			fmt.Printf("Balanced Updated!Current balance is: %.2f Tk\n", accountBalance)
		} else if choice == 3 {
			fmt.Print(("Enter your Cash out amount: "))
			var cashOutAmount float64
			fmt.Scan((&cashOutAmount))

			if cashOutAmount <= 0 {
				fmt.Println("Invalid amount! Please enter a amount greater than 0.")
			}

			if cashOutAmount > accountBalance {
				fmt.Println("Insufficient balance!")
				return
			}

			accountBalance = accountBalance - cashOutAmount
			fmt.Printf("Balanced Updated!Current balance is: %.2f Tk\n", accountBalance)
		} else {
			fmt.Println("Exiting Fake-Bikash App")

		}
	}

}

```


5. Add infinite time to run the loop and add condition for loop break

```go

// Run: go run bikash.go
package main

import "fmt"

func main() {

	accountBalance := 1000.00;

	fmt.Println("Welcome to Fake-Bikash App")
	fmt.Println("-------------------------")

	for  {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Your Balance")
		fmt.Println("2. Cash In")
		fmt.Println("3. Cash Out")
		fmt.Println("4. Exit Fake-Bikash App")

		var choice int
		fmt.Print("Enter your option: ")
		fmt.Scan(&choice)
		// fmt.Println("You have selected option:", choice)

		if choice == 1 {
			fmt.Printf("You Balance is : %.2f Tk\n", accountBalance)
		} else if choice == 2 {

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
		} else if choice == 3 {
			fmt.Print(("Enter your Cash out amount: "))
			var cashOutAmount float64
			fmt.Scan((&cashOutAmount))

			if cashOutAmount <= 0 {
				fmt.Println("Invalid amount! Please enter a amount greater than 0.")
			}

			if cashOutAmount > accountBalance {
				fmt.Println("Insufficient balance!")
				return
			}

			accountBalance = accountBalance - cashOutAmount
			fmt.Printf("Balanced Updated!Current balance is: %.2f Tk\n", accountBalance)
		} else {
			fmt.Println("-----Thanks for chosing Fake-Bikash App----")
			break;

		}
	}

}

```
6. Add switch case 


```go
// Run: go run bikash.go
package main

import "fmt"

func main() {

	accountBalance := 1000.00

	fmt.Println("Welcome to Fake-Bikash App")
	fmt.Println("-------------------------")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Your Balance")
		fmt.Println("2. Cash In")
		fmt.Println("3. Cash Out")
		fmt.Println("4. Exit Fake-Bikash App")

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




