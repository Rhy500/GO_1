package main

import (
	"fmt"
)

func main() {
	var plat string
	var platBandung bool
	fmt.Println("masukan sebuah huruf;")
	fmt.Scan(&plat)

	platBandung = plat == "D"
	fmt.Println(platBandung)
}
