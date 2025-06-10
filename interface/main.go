package main

import "fmt"

type Client struct {
	Name    string
	age     int
	active  bool
	Address Address
}

type Address struct {
	City    string
	State   string
	ZipCode int
	Address string
}

type Person interface {
	Desactive()
}

func (c Client) Desactive() {
	c.active = false
	fmt.Println("Client desactivated %v", c.active)
}

func desactivate(person Person) {
	person.Desactive()
}

func main() {

	client := Client{
		Name:   "John Doe",
		age:    30,
		active: true,
		Address: Address{
			City:    "São Paulo",
			State:   "Sp",
			ZipCode: 12345678,
			Address: "rua das flores, 123",
		},
	}

	desactivate(client)

}
