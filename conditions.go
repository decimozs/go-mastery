package main

import "fmt"

func main(){
	fmt.Println("Conditions")

	x := 10
	y := 20

	// if statement
	if x < y {
		fmt.Println("x is less than y")
	}

	// if else statement
	if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is greater than y")
	}

	// else if statement
	if x < y {
		fmt.Println("x is less than y")
	} else if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x is equal to y")
	}

	// nested if statement 
	if x > 5 {
		fmt.Println("x is greater than 5")
		if x > 9 {
			fmt.Println("x is also greater than 9")
		}
	} else {
		fmt.Println("x is less than 10")
	}
}