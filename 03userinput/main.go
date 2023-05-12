package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating of camera: ")

	//comma ok || err ok || _ ok
	input, _ := reader.ReadString('\n')

	fmt.Println("The rating of camera is: ", input)

	fmt.Printf("Type of this rating is %T", input)
}
