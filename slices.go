package main 

import "fmt"

func main(){
	fmt.Println("Slice")
	/*
		Slice are just like an array because it can also stored multiple values
		but slice are just bit more flexible and powerful in array because slice
		can shrink and grow (Dynamic Array) as you increase or decrease the value 
		in the slice.
	*/

	// in this slice we initialize the length, capacity and the value inside of the slice
	slice1 := [] int {1,2,3,4,5} 						// sliceName := []datatype{values}
	// print the properties of the slice 1
	fmt.Println("Slice 1")
	fmt.Println("Length: ", len(slice1))			// length of the slice 1
	fmt.Println("Capacity: ", cap(slice1))			// capacity of the slice 1
	fmt.Println("Value: ", slice1)					// value inside of slice 1

	// in this slice we do not initialize the length, capacity and the value inside of the slice it means that it is empty
	slice2 := [] int {} 									// sliceName := []datatype{values}
	// print the properties of the slice 2
	fmt.Println("\nSlice 2")
	fmt.Println("Length: ", len(slice2))			// length of the slice 2
	fmt.Println("Capacity: ", cap(slice2))			// capacity of the slice 2
	fmt.Println("Value: ", slice2)					// value inside of slice 2

	// creating a slice from an array
	fmt.Println("\nSlice [Slice that created with an array]")
	array := [5] int {1,2,3,4,5}
	slice := array[1:2]									// sliceName := arrayName[start:end]
	// print the properties of the slice that created with a array
	fmt.Printf("slice = %v\n", slice)
  	fmt.Printf("length = %d\n", len(slice))
  	fmt.Printf("capacity = %d\n", cap(slice))

	fmt.Println("\nSlice [Slice that created with make function]")
	// creating a slice with make function
	// with length and capacity
	fmt.Println("Slice with length and capacity")
	makeSlice := make([] int, 5,10)					// sliceName := make([] datatype, length, capacity)
	fmt.Printf("slice = %v\n", makeSlice)
  	fmt.Printf("length = %d\n", len(makeSlice))
  	fmt.Printf("capacity = %d\n", cap(makeSlice))
	// with length but omitted capacity
	fmt.Println("Slice with length and omitted capacity")
	makeSlice1 := make([] int, 5,10)					// sliceName := make([] datatype, length, capacity)
	fmt.Printf("slice = %v\n", makeSlice1)
  	fmt.Printf("length = %d\n", len(makeSlice1))
  	fmt.Printf("capacity = %d\n", cap(makeSlice1))

	// accesing elements of a slice
	fmt.Println("\nAccesing elements of a slice (Slice 1)")
	fmt.Println(slice1[0])
	fmt.Println(slice1[1])

	// chaning elements of a slice
	fmt.Println("\nChanging elements of a slice (Slice 1)")
	// change the value of the index 1 in the slice1 to 10
	slice[0] = 10
	fmt.Println("After changing the elements in the slice")
	fmt.Println(slice1[0])
	fmt.Println(slice1[1])

	// appending elements to a slice
	fmt.Println("\nAppending elements to a slice")
	// create a slice
	slice3 := [] int {1,2,3,4,5}
	// print the properties of the slice3
	fmt.Println("No appending of elements happening")
	fmt.Printf("slice3 = %v\n", slice3)
  	fmt.Printf("length = %d\n", len(slice3))
  	fmt.Printf("capacity = %d\n", cap(slice3))
	// append a value to the slice3
	fmt.Println("After appending of elements")
	slice3 = append(slice3, 6,7,8,9,10)				// sliceName = append(sliceName, appendedValue...)
	fmt.Printf("myslice1 = %v\n", slice3)
  	fmt.Printf("length = %d\n", len(slice3))
  	fmt.Printf("capacity = %d\n", cap(slice3))

	// appending one slice to another slice
	fmt.Println("\nAppending one slice to another slice")
	// two slices 
	slice4 := [] int {1,2,3,4,5}
	slice5 := [] int {6,7,8,9,10}
	// append those slices (Combine the two slices into one)
	appendSlice := append(slice4, slice5...)				// sliceName := append(slice1, slice2...) 
	fmt.Println("After appending the two slices")
	fmt.Printf("appendSlice=%v\n", appendSlice)
  	fmt.Printf("length=%d\n", len(appendSlice))
  	fmt.Printf("capacity=%d\n", cap(appendSlice))

	/*
		In terms of changing the length of the slice there are ways to change it
		[1] by re-slicing the array
		[2] by appending the items to the slice
	*/

	// memory efficiency when it comes to using slices
	// by using the copy functions
	// copy(dest, src)
	// original slice
	fmt.Println("\nMemory efficiency [Using copy function - copy()]")
	fmt.Println("Original Slice")
	numbers := [] int {1,2,3,4,5,6,7,8,9,10}
	fmt.Printf("numbers = %v\n", numbers)
  	fmt.Printf("length = %d\n", len(numbers))
  	fmt.Printf("capacity = %d\n", cap(numbers))
	// using copy function for only needed numbers
	fmt.Println("Copying the needed values from the numbers slice")
	neededNumbers := numbers[:len(numbers)-5]						// create a slice for storing the needed values from the original slice and decrease the value of it by 5 from the end
	numbersCopy := make([]int, len(neededNumbers))				// make a new slice for storing the extracted value from the neededValues slice
	copy(numbersCopy, neededNumbers)									// copy the value of the second parameters to the first paramenter 
	// print the properties after copying the needed values
	fmt.Printf("numbersCopy = %v\n", numbersCopy)
  	fmt.Printf("length = %d\n", len(numbersCopy))
  	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}