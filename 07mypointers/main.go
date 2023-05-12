package main

import "fmt"

func main() {
	fmt.Println("Welcome to learn pointers")

	//declaring a pointer using astretic operator
	var ptr *int

	//default value of pointer will be <nil> is something is not assigned
	fmt.Println("Value of pointer is ", ptr)

	var number int = 10

	// here we are creating a pointer which is refrencing to a memory which can be done by using & operator before any varible
	// below is reference of direct memory location
	// & means memory address
	// * means value of that memory address
	var refrencingPtr = &number
	fmt.Println("Value of refrencing pointer is ", refrencingPtr)  // 0xc0001a8000
	fmt.Println("Value of refrencing pointer is ", *refrencingPtr) // 10

	*refrencingPtr = *refrencingPtr + 2
	fmt.Println("Value of number is ", number) // 12 because we are adding 2 to the value of number directly using the memory location
}
