package main

import "fmt"

func main() {
	/*kamus = x, y int and hasil bool . kali3 int
	aligoritma = input(x,y) . kali3 <- x*x*x . hasil <- kali3 == y . output(hasil) */

	var x, y int
	var hasil bool

	fmt.Scan(&x)
	fmt.Scan(&y)

	hasil = y%x*x*x == 0
	fmt.Println(hasil)
}
