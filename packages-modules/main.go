package main

import (
	"curso-go/mathematics"
	"fmt"
)

func main() {

	sum := mathematics.Sum[int](10, 20)

	fmt.Println("The result of sum is", sum)

	fmt.Println(mathematics.A)

	car := mathematics.Car{
		Brand: "Ford",
		Model: "Mustang",
	}

	println("Car brand is ", car.Brand, " and model is ", car.Model)
}
