package main

import "fmt"

func main() {
	var day, month, year, maxDay int
	var tahunKabisat bool

	fmt.Scan(&day, &month, &year)

	// Validasi tahun
	tahunKabisat = (year%4 == 0 && year%100 != 0) || (year%400 == 0)

	// Validasi tanggal berdasarkan bulan
	maxDay = 0
	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		maxDay = 31
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		maxDay = 30
	} else if month == 2 {
		if tahunKabisat {
			maxDay = 29
		} else {
			maxDay = 28
		}
	}

	// Output hasil validasi
	if day >= 1 && day <= maxDay {
		fmt.Println("Tanggal Valid")
	} else {
		fmt.Println("Tanggal Tidak Valid")
	}

	if month >= 1 && month <= 12 {
		fmt.Println("Bulan Valid")
	} else {
		fmt.Println("Bulan Tidak Valid")
	}

	if year > 0 {
		fmt.Println("Tahun Valid")
	} else {
		fmt.Println("Tahun Tidak Valid")
	}
}
