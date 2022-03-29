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

func Notify(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Число не поддерживается")
	}

	s, ok := i.(Sender)
	if !ok {
		fmt.Println("Ошибка утверждения интерфейса")
		return
	}

	err := s.Send("Сообщение пустого интерфейса")
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println("Success")
}

func main() {
	email := &Email{"vladimir@gmail.com"}
	Notify(email)

	phone := &Phone{Number: 7777, Balance: 10}
	Notify(phone)

	Notify(3)
	Notify("String")

	sl := [5]int64{1, 2, 3, 4}
	Notify(sl)
}
