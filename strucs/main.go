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

func (c Client) Desactive() {
	c.active = false
	fmt.Println("Client desactivated %v", c.active)
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

	client.Desactive()

	fmt.Printf("Name: %v, Age: %v , Active: %v \n", client.Name, client.age, client.active)
	fmt.Printf("City: %v, State: %v, ZipCode: %v, Address: %v", client.Address.City, client.Address.State, client.Address.ZipCode, client.Address.Address)

}
