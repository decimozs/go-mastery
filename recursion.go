package main

import "fmt"

func main(){
	// recursion is a recursive call of itselfs and need to reach a conditions or statements to stop the process
	fmt.Println("Recursion")
	// call the function and invoke it repeatedly
	// [note] that the function will call itself if the function has no base condition the process will not stop
	increment(1)
}

// increment function where it takes a integer of x 
func increment(x int) int {
	// the base case of the function
	if x == 11 {
		// if the x is equal to the 11 it will return 0 and stop the process
		return 0
	}
	// print the corresponding value after it met the condition
	fmt.Println(x)
	// continue to print and increment the assign value until it meet the base condition
	return increment(x + 1)
}