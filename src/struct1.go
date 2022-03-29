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
	Img  *Avatar
}

func main() {
	i := new(int64)
	_ = i

	client := Client{}
	client.Img = new(Avatar)

	updateAvatar(&client)

	fmt.Printf("%#v\n", client)
}

func updateAvatar(client *Client) {
	client.Img.Url = "http://update_url"
}
