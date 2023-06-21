package main

import "fmt"

func main(){
	fmt.Println("Struct")

	// assigning the Person struct in the variable person1 and person2
	var person1 Person
	var person2 Person

	// using the variable person1 to assign a initial value to the propeties of Person struct
	person1.name = "Decimo"
	person1.age = 20
	person1.job = "Student"
	person1.salary = 500

	// using the variable person2 to assign a initial value to the propeties of Person struct
	person2.name = "Martin"
	person2.age = 20
	person2.job = "Software Engineer"
	person2.salary = 32000

	// print the assign value
	fmt.Println("Person 1")
	fmt.Println("Name: ", person1.name)
	fmt.Println("Age: ", person1.age)
	fmt.Println("Job: ", person1.job)
	fmt.Println("Salary: ", person1.salary)

	fmt.Println("\nPerson 2")
	fmt.Println("Name: ", person2.name)
	fmt.Println("Age: ", person2.age)
	fmt.Println("Job:", person2.job)
	fmt.Println("Salary: ", person2.salary)

	// [alternatives] call the function and pass the variable person1 or person2 to print its properties
	fmt.Println("\nPerson 1 [Struct as a function arguments]")
	printPerson(person1)
	fmt.Println("\nPerson 2 [Struct as a function arguments]")
	printPerson(person2)
}

// struct is a collection of data that has different data types into a single variable
type Person struct {
	name string			// string
	age int				// int
	job string			// string
	salary int			// int
}

// you can pass a struct as a function arguments
func printPerson(pers Person){
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job:", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
