package main

func main() {
	a := 1
	b := 2
	c := 3
	if a > b {
		println(a)
	} else {
		print(b)
	}

	if a > b && b > c {
		println("And")
	}

	if a > b || b > c {
		println("Or")
	}

	switch a {
	case 1:
		println("A")
	case 2:
		println("B")
	case 3:
		println("C")
	default:
		println("Default")
	}
}
