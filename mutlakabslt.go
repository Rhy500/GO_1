package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x <= 0 {
		fmt.Println(x * -1)
	} else {
		fmt.Println(x)
	}
}
