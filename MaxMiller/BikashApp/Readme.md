🔲 Run: 
- `cd BikashApp` 
- `go run bikash.go` 

I want to make a simple **bikash app** which we see in our real life
by typing `*247#`.

<b>Step by setp how i devlop the app.

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
	fmt.Println("You have selected option:", choice)

	if choice == 1 {
		fmt.Printf("You Balance is : %.2f Tk\n", accountBalance);
	}else if choice == 2 {
		fmt.Print("Enter your Cash in amount: ");
		var cashInAmount float64;
		fmt.Scan(&cashInAmount)

		accountBalance=accountBalance+cashInAmount;
		fmt.Printf("Your Current balance is: %.2f Tk\n", accountBalance);
	}
}
```
