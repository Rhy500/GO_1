package main

import "fmt"

func main() {
	var x, digit, total int
	fmt.Scan(&x)
	total = 0
	for x > 0 {
		digit = x % 10
		total = total + digit
		fmt.Print(digit, " ")
		x = x / 10
	}
	fmt.Println(" ")
	fmt.Println(total)

}
