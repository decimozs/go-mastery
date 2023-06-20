package main

import "fmt"

func main(){
	fmt.Println("Switch Statement")

	month := 1

	// switch the expresstion you want to validate which is month
	switch  month {
	case 1:
		fmt.Println("January")			// code  block
	case 2:
		fmt.Println("February")			// code  block
	case 3:
		fmt.Println("March")				// code  block	
	case 4:
		fmt.Println("April")				// code  block
	case 5:
		fmt.Println("May")				// code  block
	case 6:
		fmt.Println("June")				// code  block
	case 7:
		fmt.Println("July")				// code  block
	case 8:
		fmt.Println("August")			// code  block
	case 9:
		fmt.Println("September")		// code  block
	case 10:	
		fmt.Println("October")			// code  block
	case 11:
		fmt.Println("November")			// code  block
	case 12:
		fmt.Println("December")			// code  block
	default : 
		fmt.Println("Invalid no of month")	// code  block if the case is not met
	}

}