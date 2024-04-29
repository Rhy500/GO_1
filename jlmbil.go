package main

import "fmt"

func main() {
	var bil, hasil int
	fmt.Scan(&bil)
	hasil = 0
	for bil != 0 {
		hasil = hasil + bil
	}
	fmt.Println(hasil)
}
