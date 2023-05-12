package main

import (
	"fmt"
	"sort"
)

func main() {
	//slices is an data type in go lang
	//slices are actually an array and it has some abstraction layes and more features
	//slices provide us a dynamic array

	//when we dont declare a size that means we are using slices and which menas we need to initialize it
	var fruits = []string{"Apple", "Orange", "Banana"}
	fmt.Printf("Data type of fruits is %T\n", fruits)

	fruits = append(fruits, "Mango", "Grapes") //appending the data in fruits variable and it takes the first param as array/slice and second value is actaul data
	fmt.Println(fruits)

	//it actaully slice the data from the array by 2 elements
	//it takes 1:3 which means slice from index 1 to 3 and 4 will be non included
	fruits = append(fruits[1:3], "peach")
	fmt.Println(fruits) //[Orange Banana]

	//so the size of array is 4
	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 345
	highscores[2] = 465
	highscores[3] = 867
	// highscores[4] = 86  //error

	//this will append the data but size is 4
	//but the append method will reallocate the size of slices
	highscores = append(highscores, 56, 12, 132, 123)
	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println("Highscores sorted:", highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	//how to remove the element from the slices based on index
	var courses = []string{"HTML", "CSS", "JS", "React", "Node", "Golang"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //append method will reallocate the memory
	fmt.Println(courses)

}
