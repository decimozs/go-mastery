package main

import "fmt"

func main(){
	// range
	/*
		for index, value := (it can be an array or slice or map) {
			statement or expression 
		}
	*/
	fmt.Println("Range")
	// declare a array of strings of cars 
	cars := [3] string {"Nissan Silvia S15", "Mazda RX7", "Evo Lancer"}
	for index, value := range cars{
		fmt.Println("\n",index, value)
	}
}