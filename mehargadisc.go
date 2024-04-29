package main

import "fmt"

func main() {
	var harga_asli, harga_discon float64
	var discon float64
	fmt.Scan(&harga_asli, &discon)
	harga_discon = (harga_asli * (1 - discon))
	fmt.Println(harga_discon)
}
