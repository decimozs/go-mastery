package main 

import "fmt"

func main(){
	// assigning multiple values
	var a,b,c,d,e int = 1,2,3,4,5

	// assigning multiples values in a block way
	var(
		t int = 1
		v int = 2
		x int = 3
		y int = 4
		z int = 5
	)

	fmt.Println("Multiple Values")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println("Multiple Values in a Block Way")
	fmt.Println(t)
	fmt.Println(v)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}