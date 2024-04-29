package main

import "fmt"

func main() {
	var banyak, i int
	var kata string
	//givar k string
	fmt.Scan(&kata, &banyak)
	for i = 1; i <= banyak; i++ {
		fmt.Println(kata)
	}
}
