package main

import "fmt"

func main() {
	var p, jumlah int
	fmt.Scan(&p)
	if p <= 3*5 {
		jumlah = p / 5
		if (p % 5) != 0 {
			jumlah = jumlah + 1
		}
		fmt.Println("Dewasa ", jumlah)
	} else if p > 3*5 && p <= (3*5+5*2) {
		jumlah = (p - 3*5) / 2
		if ((p - 3*5) % 2) != 0 {
			jumlah = jumlah + 1
		}
		fmt.Println("Dewasa 3, kecil ", jumlah)
	} else {
		jumlah = p - (3*5 + 5*2)
		fmt.Println("Dewasa 3, kecil 5, dan ", jumlah, " tak berangkat")
	}
}
