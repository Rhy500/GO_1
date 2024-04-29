package main

import "fmt"

func main() {
	// kamus = x, y ; int and hasil ; bool
	// algoritma = input(x,y) . hasil <- ( y mod x) == 0 . output (hasil)
	var x, y int
	var hasil bool
	fmt.Scan(&x)
	fmt.Scan(&y)

	hasil = y%x == 0
	fmt.Println(hasil)

}
