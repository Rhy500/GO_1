package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input rune
	fmt.Println("masukan sebuah huruf;")
	_, err := fmt.Scanf("%c", &input)
	if err != nil {
		fmt.Println("input tidak valid.")
		return
	}

	isCapital := unicode.IsUpper(input)

	fmt.Printf("Output: %v\n", isCapital)

}
