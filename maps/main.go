package main

import "fmt"

func main() {
	salarios := map[string]int{"Matheus": 10000, "joão": 5000, "Carlos": 7000}

	fmt.Println(salarios["Matheus"])
	delete(salarios, "Matheus")

	fmt.Println(salarios["Matheus"])
	salarios["Matias"] = 1200000
	fmt.Println(salarios["Matias"])

	sal := make(map[string]int)
	sal1 := map[string]int{}

	sal["Fabio"] = 100
	sal1["Fabia"] = 200

	fmt.Printf("O salario de Fabia é %v \n", sal1["Fabia"])

	for nome, valor := range salarios {
		fmt.Printf("O salario de %s é %d \n", nome, valor)
	}
	for _, valor := range salarios {
		fmt.Printf("O salario é %d \n", valor)
	}

}
