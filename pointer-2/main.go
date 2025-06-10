package main

import "fmt"

func main() {
	var v1 int = 10
	var v2 int = 20

	fmt.Println("Sum:", sum(&v1, &v2))
	fmt.Println("Value of v1:", v1)
}

func sum(a, b *int) int {
	*a = 22
	return *a + *b
}
