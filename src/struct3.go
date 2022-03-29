package main

import (
	"client"
	"fmt"
)

func main() {
	avatar := client.Avatar{
		Url: "url",
	}

	c := client.NewClient("John", 27, avatar)
	fmt.Println(c)
}
