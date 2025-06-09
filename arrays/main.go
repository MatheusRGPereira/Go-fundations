package main

import "fmt"

func main() {
	var firstArray [3]int

	firstArray[0] = 1
	firstArray[1] = 2
	firstArray[2] = 3

	for i, v := range firstArray {
		fmt.Printf("o meu indice é %d e o valor é %d \n", i, v)
	}

}
