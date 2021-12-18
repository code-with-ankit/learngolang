package main

import "fmt"

type User struct {
	Name  string
	Email string
	age   int
}

func main() {
	rob := User{"rob", "rob@gmail.com", 34}
	fmt.Printf("%v\n", rob)
	fmt.Printf("%+v\n", rob)
	fmt.Printf("%v\n", rob.Email)

	var sam = new(User)
	sam.Name = "Ankit"
	sam.age = 25
	sam.Email = "ankit@gmail.com"
	fmt.Printf("%+v\n", sam)

	var tobby = &User{"rahul", "rahul@gmail.com", 25}
	fmt.Printf("%+v\n", tobby)
}
