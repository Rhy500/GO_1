package main

import "fmt"

func main() {
	var x, i int

	fmt.Scan(&x)
	for i = 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Print(i, " ")
		}
		if i*1 == x/1 || x != 1 {
			fmt.Println("true")
		} else {
			fmt.Println("true")
		}
	}

}
