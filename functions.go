package main

import "fmt"

func main(){
	fmt.Println("Functions")
	label()									// calling the function and invoking it
	fmt.Println(sum(5,5))				// calling the function and printing the sum
	fmt.Println(difference(5,5))		// calling the function and printing the difference
	fmt.Println(product(5,5))			// calling the function and printing the product
	fmt.Println(quotient(5,5))			// calling the function and printing the quotient
}

// label function
func label(){
	// return the print statement
	fmt.Println("Result")
}

// sum function, take a parameters of a integer x and y and returning the sum of it by adding them (x + y)
func sum(x int, y int) int {
	return x + y
}

// named return values
// difference function, take also a parameters of a integer x and y and returning the difference of it in difference variable by subtracting them (x - y)
func difference(x int, y int) int {
	difference := x - y
	return difference
}

func product(x int, y int) int {
	return x * y
}

func quotient(x int, y int) int {
	return x / y
}