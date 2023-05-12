package main

import "fmt"

// global declaration
// constant declaration
// public variable
// if varible has first charcter capital then it is public variable according to go lang. Instead of any public keyword.
// This variable is accesssible by any other file from this folder or this file.
const Accesstoken string = "echgmwoldalkm"

func main() {
	//Type string
	var username string = "viraj.doshi"
	fmt.Println("username: ", username)
	fmt.Printf("variable is of type: %T \n\n", username)

	//Type boolean
	var isActive bool = true
	fmt.Println("isActive: ", isActive)
	fmt.Printf("variable is of type: %T \n\n", isActive)

	//Type unit8 range will be 0 - 255 (it also has alias unit, unit16, unit32, unit64)
	var smallVal uint8 = 255
	fmt.Println("smallVal: ", smallVal)
	fmt.Printf("variable is of type: %T \n\n", smallVal)

	//Type float32 range will be upto 32 bits and after decimal it will show upto 5 numbers.
	var smallFloat float32 = 255.6473284324983985594
	fmt.Println("smallFloat: ", smallFloat)
	fmt.Printf("variable is of type: %T \n\n", smallFloat)

	//Type float64 range will be upto 64 bits and after decimal it will show upto 13 numbers.
	var bigFloat float64 = 255.6473284324983985594
	fmt.Println("bigFloat: ", bigFloat)
	fmt.Printf("variable is of type: %T \n\n", bigFloat)

	//Type int also a alias of unit and default value will be 0
	var anothervariable int
	fmt.Println("anotherVariable: ", anothervariable)
	fmt.Printf("variable is of type: %T \n\n", anothervariable)

	//implicit type
	//also if we dont declare variable type then lexer will decide it type and then you wont change the type in future
	var website = "Google.com"
	fmt.Println("website: ", website)
	fmt.Printf("variable is of type: %T \n\n", website)

	//no var style or walrus operator will allow the user to declare variable without var keyword but you need to declare it by ":=".
	//Walrus operator is only allowed inside a function or a method
	numberOfUsers := 1500.0
	fmt.Println("numberOfUsers: ", numberOfUsers)
	fmt.Printf("variable is of type: %T \n\n", numberOfUsers)

	//using public variable, global variable
	fmt.Println("Accesstoken: ", Accesstoken)
	fmt.Printf("variable is of type: %T \n\n", Accesstoken)
}
