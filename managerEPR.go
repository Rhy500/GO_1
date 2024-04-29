package main

import "fmt"

func main() {
	var b1, b2, b3, b4, b5 string
	fmt.Scan(&b1, &b2, &b3, &b4, &b5)

	if (b1 == "kalah") && (b2 == "kalah") && (b3 == "kalah") && (b4 == "kalah") && (b5 == "kalah") {
		fmt.Println("dipecat")
	} else {
		fmt.Println("tidak dipecat")
	}
}
