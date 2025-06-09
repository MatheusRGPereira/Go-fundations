package main

import (
	"errors"
	"fmt"
)

func main() {

	val, error := sumError(1, 2)
	if error != nil {
		fmt.Println("Erro:", error)
	}
	fmt.Println(val)

	fmt.Println(sum(100, 2))

}

func sum(a, b int) (int, bool) {
	if a+b > 10 {
		return a + b, true
	}
	return a + b, false
}

func sumError(a, b int) (int, error) {
	if a+b > 10 {
		return a + b, errors.New(("the sum is greater than 10"))
	}
	return a + b, nil
}
