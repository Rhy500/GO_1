package main

import "fmt"

func main() {
	var n, total, digit int
	fmt.Scanln(&n)
	total = 0

	for n > 0 {
		digit = n % 10
		total = total + digit
		fmt.Print(digit, " ")
		n = n / 10
	}
	fmt.Println(" ")
	fmt.Println(total)

}
