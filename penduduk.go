package main

import "fmt"

func main() {
	var jumlahPenduduk int
	var pedkawa, kelahiran, imigrasi, kematian, emigrasi int

	fmt.Scan(&pedkawa)
	fmt.Scan(&kelahiran)
	fmt.Scan(&imigrasi)
	fmt.Scan(&kematian)
	fmt.Scan(&emigrasi)

	jumlahPenduduk = (pedkawa + kelahiran + imigrasi - kematian - emigrasi)
	fmt.Println(jumlahPenduduk)
}
