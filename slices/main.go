package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("len=%d cap=%d %v \n", len(slice), cap(slice), slice)

	fmt.Printf("len=%d cap=%d %v \n", len(slice[:0]), cap(slice[:0]), slice[:0])

	fmt.Printf("len=%d cap=%d %v \n", len(slice[:4]), cap(slice[:4]), slice[:4])

	fmt.Printf("len=%d cap=%d %v \n", len(slice[2:]), cap(slice[2:]), slice[2:])

	slice = append(slice, 11)
	fmt.Printf("len=%d cap=%d %v \n", len(slice[:2]), cap(slice[:2]), slice[:2])
}
