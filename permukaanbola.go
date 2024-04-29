package main

import "fmt"

func main() {
	var digit, doubled float64
	fmt.Println("you number")
	fmt.Scan(&digit)

	doubled = (digit*digit + (1.0 / digit)) * (digit*digit + (1.0 / digit))
	fmt.Println(doubled)
}
