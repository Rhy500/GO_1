package main

import "fmt"

func main() {
	var b1, b2, b3, b4 int
	fmt.Scan(&b1, &b2, &b3, &b4)
	if b1 > b2 && b1 < b3 && b1 > b4 && b2 < b3 && b2 > b4 && b3 > b4 {
		fmt.Println(b3, b4)
	} else if b1 == b2 && b2 == b3 && b3 == b4 {
		fmt.Println(b1, b4)
	} else {
		fmt.Println(b2, b4)
	}
}
