package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"} // array/slices of days
	fmt.Println("Days of the week:", days)

	//simple for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//for loop using range
	for i := range days {
		fmt.Println(days[i])
	}

	//for loop with index and value
	for index, day := range days {
		fmt.Printf("The index of %v is %v\n", day, index)
	}

	rougueValue := 1

	for rougueValue < 10 {

		//goto statement
		if rougueValue == 2 {
			goto lco
		}

		//continue statement
		if rougueValue == 5 {
			rougueValue++
			continue
		}

		//break statement
		if rougueValue == 9 {
			break
		}

		fmt.Println("value is: ", rougueValue)
		rougueValue++
	}

	//label declared for goto statement
lco:
	fmt.Println("Jumping on line 48 beacause of goto statement")
}
