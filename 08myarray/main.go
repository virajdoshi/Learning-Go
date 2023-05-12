package main

import "fmt"

func main() {

	//can declare array by provideing []type and in array brackets we need to provide an number to define the size of array
	var friuts [4]string

	friuts[0] = "Apple"
	friuts[1] = "Orange"
	friuts[3] = "Banana"

	fmt.Println("Fruits:", friuts)                   //[Apple Orange  Banana] in this we can see the big space between the elements which means we have not filed all the values upto size 4
	fmt.Println("Length of fruits is ", len(friuts)) //length of furits is 4 but values is 3 so it does not give the actaul size of array

	//other way to declare array
	var vegitables = [3]string{"Brocolli", "Potato", "Carrot"}
	fmt.Println("Vegetables:", vegitables)

	//len(varible) will give the fixed size of array
}
