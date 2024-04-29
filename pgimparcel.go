package main

import "fmt"

func main() {
	var beratParcelGr, beratKg, beratSisaGr, biayaPerKg, tambahanBiaya, totalBiaya int
	fmt.Scan(&beratParcelGr)
	beratKg = beratParcelGr / 1000
	beratSisaGr = beratParcelGr % 1000
	biayaPerKg = 10000
	if beratKg <= 10 {
		if beratSisaGr >= 500 {
			tambahanBiaya = beratSisaGr * 5
		} else {
			tambahanBiaya = beratSisaGr * 15
		}
	}
	totalBiaya = (beratKg * biayaPerKg) + tambahanBiaya
	//fmt.Println(beratKg, beratSisaGr)
	//fmt.Println("%d kg + %d gr\n", beratKg, beratSisaGr)
	//fmt.Println("Rp. %d + Rp. %d = Rp. %d\n", beratKg*biayaPerKg, tambahanBiaya, totalBiaya)
	fmt.Println(beratKg*biayaPerKg, tambahanBiaya, "=", totalBiaya)
}
