package main

import "fmt"

func main() {
	var total int
	var bersedia, kartu, diskon, cashback bool
	fmt.Scan(&total, &bersedia)
	kartu = bersedia
	diskon = (total >= 100000)
	cashback = (total >= 200000) && (kartu == true)
	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
}
