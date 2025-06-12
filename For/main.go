package main

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove", "dez"}
	for k, numero := range numeros {
		println(k, numero)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}
}
