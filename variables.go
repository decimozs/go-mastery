package main

import "fmt"

func main(){
	// Explicit declaration of variables
	// integer - explicit declaration
	var explicitNumber int = 1;
	// float - explicit declaration
	var explicitFloat float32 = 10.0;
	// string - explicit declaration
	var explicitName string = "Decimo"
	// boolean - explicit declaration
	var explicitCheck bool = true

	// Inferred declaration of variables
	// integer - inferred declaration
	inferredNumber := 2
	// float - inferred declaration
	inferredNumberFloat := 2.0
	// string - inferred declaration
	inferredName := "Dez"
	// boolean - inferred declaration
	inferredCheck := false;
	
	fmt.Println("Explicit Declaration of Variables");
	fmt.Println(explicitNumber)
	fmt.Println(explicitFloat)
	fmt.Println(explicitName)
	fmt.Println(explicitCheck)

	fmt.Println("Inferred Declaration of Variables");
	fmt.Println(inferredNumber)
	fmt.Println(inferredNumberFloat)
	fmt.Println(inferredName)
	fmt.Println(inferredCheck)
}