package main

import "fmt"

func main() {
	var jam, menit, detik, waktu int
	fmt.Scan(&waktu)

	jam = waktu / 3600
	menit = waktu % 3600 / 60
	detik = waktu % 3600 % 60

	fmt.Println(jam, menit, detik)
}
