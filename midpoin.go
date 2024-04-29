package main

import "fmt"

func main() {
	var a, b, c int
	// a hasil nilai tengahnya , b = batas atas , c=batas bawah
	var nilaitengah int
	fmt.Scan(&a, &b, &c)

	nilaitengah = (b + c) / 2
	if nilaitengah == a {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
