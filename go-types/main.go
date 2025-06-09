package main

import "fmt"

type Id int

var (
	b bool    = true
	c int     = 1
	d string  = "String-Type"
	e float64 = 33.2
	f Id      = 7
)

func main() {

	fmt.Printf("O tipo da variavel e é %T", e)
}
