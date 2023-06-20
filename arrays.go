package main

import "fmt"

func main(){
	// arrays without specific length
	// declaring arrays in a explicit way with inferred length
	var explicitArray = [...] int {1,2,3,4,5,6,7,8,9,10}
	// declaring arrays in a inferred way with inferred length
	inferredArray := [...] int {1,2,3,4,5,6,7,8,9,10}

	// arrays with specific length
	// declaring arrays with a specific length in a explicit way
	var explicitStringArray = [10] string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	// declaring arrays with a specific length in a inferred way
	inferredStringArray := [10] string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	fmt.Println("Arrays")
	fmt.Println(explicitArray)
	fmt.Println(inferredArray)
	fmt.Println(inferredStringArray)
	fmt.Println(explicitStringArray)

	fmt.Println("\nAccesing elements in array")
	fmt.Println(explicitArray[0])
	fmt.Println(inferredArray[1])
	fmt.Println(inferredStringArray[2])
	fmt.Println(explicitStringArray[3])

	// changing elements in the array
	explicitArray[0] = 2
	inferredArray[1] = 3
	inferredStringArray[2] = "d"
	explicitStringArray[3] = "f"

	fmt.Println("\nAfter changing elements in array")
	fmt.Println(explicitArray[0])
	fmt.Println(inferredArray[1])
	fmt.Println(inferredStringArray[2])
	fmt.Println(explicitStringArray[3])

	// length of the array
	fmt.Println("\nLength of the array")
	fmt.Println(len(explicitArray))
	fmt.Println(len(inferredArray))
	fmt.Println(len(inferredStringArray))
	fmt.Println(len(explicitStringArray))
}