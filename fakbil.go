package main

import "fmt"

func main() {
	var n, i int

	fmt.Scanln(&n)

	fmt.Printf("The factors of %d are: ", n)
	for i = 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d", i)
		}
	}
}
