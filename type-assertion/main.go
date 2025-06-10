package main

import "fmt"

func main() {
	var hello interface{} = "Hello, World!"
	println(hello.(string))
	res, ok := hello.(int)

	fmt.Printf("Type assertion result: %v, ok: %v\n", res, ok)

}
