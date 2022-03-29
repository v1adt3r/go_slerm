package main

import (
	"client_person"
	"fmt"
)

func main() {
	c := client_person.Client{}

	fmt.Println(c)
	c.UpdateAvatar()
	fmt.Println(c)

	c.Name = "Client name"
	c.Person.Name = "Person name"

	fmt.Println(c.GetName())
	fmt.Println(c.Person.GetName())
}
