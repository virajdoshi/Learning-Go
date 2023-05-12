package main

import "fmt"

func main() {
	var userLoginCount int = 6

	if userLoginCount > 10 {
		fmt.Println("Limit Exceeded")
	} else if userLoginCount > 5 {
		fmt.Println("otp send successfully")
	} else {
		fmt.Println("otp not send")
	}

	if 12%2 == 0 && 12%3 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 12; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}
}
