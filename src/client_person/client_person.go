package client_person

type Avatar struct {
	Url  string
	Size int64
}

type Person struct {
	Name string
	Age  int
}

type Client struct {
	Id   int64
	Img  Avatar
	Name string
	Person
}

func (p Person) GetName() string {
	return p.Name
}

func (c Client) GetName() string {
	return c.Name
}

func (c Client) HasAvatar() bool {
	return c.Img.Url != ""
}

func (c *Client) UpdateAvatar() {
	c.Img.Url = "http://newUrl.ru"
}
