package main

import "fmt"

func main() {
	var n, u int

	fmt.Scan(&n)
	u = 0
	for u != n {
		u = u + 1
		fmt.Println(u, " ")
	}
}
