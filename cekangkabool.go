package main

import "fmt"

func main() {
	var bil int
	var hasil bool

	fmt.Scan(&bil)
	hasil = bil%2 == 0
	fmt.Println(hasil)
}
