package main

import "fmt"

func main() {
	var k int

	fmt.Scanf("%c", &k)

	if k >= 'A' && k <= 'Z' {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
