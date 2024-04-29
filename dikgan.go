package main

import "fmt"

func main() {
	var k, hasil1, hasil2, hasil int
	fmt.Scan(&k)
	hasil1 = k / 10
	hasil2 = k % 10
	hasil = (hasil1 * 11 * 100) + (hasil2 * 11)

	fmt.Println(hasil)

}
