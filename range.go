package main

import "fmt"

func main(){
	fmt.Println("Range")
	// range
	/*
		for index, value := (it can be an array or slice or map) {
			statement or expression 
		}
	*/
	
	// declare a array of strings of cars 
	fmt.Println("Array")
	array1 := [3] string {"Nissan Silvia S15", "Mazda RX7", "Evo Lancer"}
	for index, value := range array1{
		fmt.Println("\n",index, value)
	}

	// declare a slice of integers
	fmt.Println("\nSlice")
	slice1 := []int{1,2,3,4,5,6,7,8,9,10}
	for index, value := range slice1{
		fmt.Println("\n", index, value)
	}

	// declare a map that has integer key and string values 
	fmt.Println("\nMap")
	map1 := map [int] string{
		1321 : "Ferrari",
		1 : "Nissan",
		12 : "Honda",
		33 : "Mazda",
	}
	for index, value := range map1{
		fmt.Println("\n", index, value)
	}
}