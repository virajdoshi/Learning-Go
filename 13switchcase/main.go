package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(7) + 1
	fmt.Println("Value of dice is ", dice)

	switch dice {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough // The keyword fallthrough allows us to execute the next case block without checking its condition
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and Roll a dice again")
	default:
		fmt.Println("Hey it is an illegal move")
	}
}
