package main

import "fmt"

func main() {
	fmt.Println("My Default main function")
	greeting()

	var result int = add(10, 20)
	fmt.Println("Result: ", result)

	proResultInt, proResultString := proAdder(10, 20, 30, 19, 20)
	fmt.Println(proResultString, proResultInt)
}

// normal function
func greeting() {
	fmt.Println("Hello from golang greeting function")
}

// function with simple arguments and return
func add(a int, b int) int {
	return a + b
}

// will receive multiple arguments and all the arguments are type of int and it returns int
// here the argumants is slice of int
// can also return multiple values and it should be wrap up in a brackets
func proAdder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "Hi my total value is "
}
