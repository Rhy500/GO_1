package main

import "fmt"

func main() {
	var nama string
	var jumlah int //bool
	var i int
	fmt.Scan(&nama, &jumlah)
	for i = 1; i <= int(jumlah); i++ {
		fmt.Println(nama)
	}
}
