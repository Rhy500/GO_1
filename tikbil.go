package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a < b && b < c && c > a {
		fmt.Println(a, b, c)
	} else if a > b && b > c && c < a {
		fmt.Println(c, b, a)
	} else {
		if a > b && b == c && c < a {
			fmt.Println(b, c, a)
		}
	}
}
