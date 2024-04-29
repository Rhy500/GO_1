package main

import "fmt"

func main() {
	var a, b, tp int
	fmt.Scan(&a, &b, &tp)
	if tp < a+b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
