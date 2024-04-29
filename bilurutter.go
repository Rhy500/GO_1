package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a < b && b < c && c > a {
		fmt.Println(c)
	} else {
		if a < b && b > c && c > a {
			fmt.Println(b)
		}
	}
}
