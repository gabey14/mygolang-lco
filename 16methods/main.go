package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// NOTE : No Inheritance in Golang & No Super or Parent

	abey := User{"Abey George", "abeygeorge14@gmail.com", true, 19}

	fmt.Println(abey)

	fmt.Printf("Abey's Details are %+v\n", abey)
	fmt.Printf("Name is %v and email is %v.\n", abey.Name, abey.Email)

	abey.GetStatus()
	abey.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", abey.Name, abey.Email)

}

// NOTE : Structs are like classes in Java and the name of the struct should be capitalized
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
