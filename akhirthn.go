package main

import "fmt"

func main() {
	var total int
	var x bool
	var kartu, Diskon, Cashback bool

	fmt.Scan(&total, &x)
	Diskon = total >= 100000
	kartu = Diskon && x
	Cashback = total >= 200000 && kartu

	if Diskon {
		total = (total - (total / 100 * 10))
		if Cashback {
			total = total - 75000
		}

	}
	fmt.Println("kartu?", kartu, " ", "Diskon?", Diskon, " ", "Cashback?", Cashback, " ", "RP.", total)
}
