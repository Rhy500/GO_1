package main

import "fmt"

func main() {
	var x, y, hasil float64
	fmt.Scan(&x, &y)
	hasil = y - x
	if hasil > 0 {
		fmt.Println("Peningkatan sebesar", hasil)
	} else if hasil < 0 {
		fmt.Println("Penurunan sebesar", hasil)
	} else {
		fmt.Println("tetap")
	}
}
