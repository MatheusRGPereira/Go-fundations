package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func sum[T int | float64](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func sum2[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func compare[T comparable](a, b T) bool {
	return a == b

}

func main() {

	m := map[string]int{"Carlos": 100, "Maria": 200, "João": 300}
	m2 := map[string]float64{"Carlos": 100.50, "Maria": 200.75, "João": 300.25}
	m3 := map[string]MyNumber{"Carlos": 100, "Maria": 200, "João": 300}

	println(sum(m))
	println(sum(m2))

	println(sum2(m))
	println(sum2(m2))

	println(sum2(m3))
	println(compare(1, 1))

}
