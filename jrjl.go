package main

import "fmt"

func main() {
	var x int
	var A, B, C, D, E int
	fmt.Scan(&x)

	//if x == A || x == E {
	//	fmt.Println("D")
	//} else if x == D {
	//	fmt.Println("E")
	//} else if x == C {
	//fmt.Println("B")
	//} else {
	//if x == B {
	//fmt.Println("C")
	//}
	//}

	for x == A || x == E {
		x = D
	}
	for x == D {
		x = E
	}
	for x == C {
		x = B
	}
	for x == B {
		x = C
	}
	fmt.Println(x)
}
