package main

import "fmt"

func main(){
	// loop
	/*
		for initialization, condition, counter {
			statement or expressions
		}
	*/
	fmt.Println("Loop")
	for i := 0; i <= 5; i++{
		fmt.Println(i)
	}

	fmt.Println("\nNested Loops")
	// declare two arrays
	array1 := [5] int {1,2,3,4,5}
	array2 := [5] string {"a","b","c","d","e"}
	// first loop or the outer loop goes to the first array
	for i := 0; i < len(array1); i++{
		// second loop or the inner loop goes to the second array
		for i := 0; i < len(array2); i++{
			fmt.Println(array1[i], array2[i])
		}
	}

	fmt.Println("\nLoop with continue statement")
	// loop with continue statement
	for i := 0; i <= 5; i++{
		// conditions meet if the i value is not divisible by 2 (odd number) then it will skip it
		if(i % 2 != 0){
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("\nLoop with break statement")
	// loop with continue statement
	for i := 0; i <= 5; i++{
		// conditions meet if the i value is equal to four (4) then it will break or stop the iteration process
		if(i == 4){
			break
		}
		fmt.Println(i)
	}
}