package main

import "fmt"

func main(){
	array1 := [5] int {} // not initialized = [0,0,0,0,0]
	array2 := [5] int {1,2,3} // partially initialized = [1,2,3,0,0]
	array3 := [5] int {1,2,3,4,5} // fully initialized = [1,2,3,4,5]

	fmt.Println("Array Initialization")
	fmt.Println(array1) 
	fmt.Println(array2) 
	fmt.Println(array3) 
}