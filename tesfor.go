package main

import "fmt"

func main() {
	var n, g int
	fmt.Scan(&n)
	g = 0
	for g != n {
		g = g + 1
		fmt.Println(g, " ")
	}
}
