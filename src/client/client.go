package client

type Avatar struct {
	Url  string
	Size int64
}

type Client struct {
	id   int64
	name string
	age  int
	img  Avatar
}

func NewClient(name string, age int, img Avatar) Client {
	return Client{
		id:   7,
		name: name,
		age:  age,
		img:  img,
	}
}

func (c Client) HasAvatar() bool {
	return c.img.Url != ""
}

func (c *Client) UpdateAvatar() {
	c.img.Url = "http://newUrl.ru"
}
