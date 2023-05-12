package main

import "fmt"

func main() {

	//below in array bracket is a key is type string and value data type is string
	languages := make(map[string]string)

	languages["js"] = "NodeJs"
	languages["py"] = "python"
	languages["go"] = "golang"

	fmt.Println("list of languages: ", languages)
	fmt.Println("Type of js is: ", languages["js"])

	//deleting the key from map
	delete(languages, "js")
	fmt.Println("list of languages: ", languages)

	//for loop to print the key and value
	for _, value := range languages {
		fmt.Printf("value is %v\n", value)
	}
}
