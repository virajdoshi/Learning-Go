package main

import "fmt"

func main() {
	//no inheritance and no super keyword or parent is in golang

	user := User{Name: "Viraj", Age: 24, Email: "virajdoshi62@ril.com", Status: true}

	//get much details of structure then we use +v to print the details
	fmt.Printf("Viraj details are: %+v\n", user)
	fmt.Printf("Name is %v and email is %v\n", user.Name, user.Email)

	user.getStatus()
	user.generateNewMail()

	fmt.Printf("Name is %v and email is %v\n", user.Name, user.Email)
}

// User is a struct where first letter is capital so it can be used by any one and it is public
// it is like a class and it can be exported and all this fields are accesible by anyone thats why first letter is capital
type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

// declaring method for the struct same we use in mongodb collection
func (u User) getStatus() {
	fmt.Println("Is user is active: ", u.Status)
}

// it actually creates a copy of email does not overwrite the old one so to do that pointer is required
func (u User) generateNewMail() {
	u.Email = u.Name + "@gmail.com"
	fmt.Println("New email is: ", u.Email)
}
