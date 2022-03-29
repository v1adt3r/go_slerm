package main

import "fmt"

type Caller interface {
	Call(number int) error
}

type IPhone interface {
	Caller
	Sender
	MyFunc()
}

type Sender interface {
	Send(msg string) error
}

type Phone struct {
	Number  int
	Balance int
}

type Email struct {
	Address string
}

func (e *Email) Send(msg string) error {
	fmt.Printf("Сообщение отправленно \"%v\" отправлено по почте на адрес %v \n", msg, e.Address)
	return nil
}

func (p *Phone) Send(msg string) error {
	fmt.Printf("Сообщение отправленно \"%v\" отправлено на телефон %v \n", msg, p.Number)
	return nil
}

func Notify(s Sender) {
	err := s.Send("Notify message")
	if err != nil {
		fmt.Println(err)
		return
	}

	switch s.(type) {
	case *Email:
		fmt.Println("Сообщение по почте")
	case *Phone:
		phone := s.(*Phone)
		fmt.Println(phone.Balance)
	}

	fmt.Println("Success")
}

func main() {
	email := &Email{"vladimir@gmail.com"}
	Notify(email)

	phone := &Phone{Number: 7777, Balance: 10}
	Notify(phone)
}
