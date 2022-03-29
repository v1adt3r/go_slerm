package main

import "fmt"

type Avatar struct {
	Url  string
	Size int64
}

type Client struct {
	Id   int64
	Name string
	Age  int
	Img  Avatar
}

func (c Client) HasAvatar() bool {
	if c.Img.Url != "" {
		return true
	}

	return false
}

func (c *Client) UpdateAvatar() {
	c.Img.Url = "http://newUrl.ru"
}

func main() {
	client := Client{
		Id:   0,
		Name: "John",
		Age:  27,
		Img: Avatar{
			Url:  "http://url.ru",
			Size: 10,
		},
	}

	fmt.Print(client)
	client.UpdateAvatar()
	fmt.Print(client)
}
