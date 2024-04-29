package main

import "fmt"

func main() {
	var bil int
	fmt.Scan(&bil)

	if bil%3 == 0 && bil%5 == 0 {
		fmt.Println("kelipatan 3 dan 5")
	} else if bil%3 == 0 {
		fmt.Println("kelipatan 3")
	} else if bil%5 == 0 {
		fmt.Println("kelipatan 5")
	} else {
		fmt.Println()
	}
}
