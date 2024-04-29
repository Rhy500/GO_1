package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan sebuah bilangan bulat positif: ")
	fmt.Scan(&input)

	// Menampilkan pola
	for i := 1; i <= input; i++ {
		// Cek apakah kita berada di tepi luar pola atau di bagian dalam
		// yang harus diisi dengan angka atau spasi
		if i == 1 || i == input {
			for j := 1; j <= input; j++ {
				fmt.Print(i)
				if j < input {
					fmt.Print(" ")
				}
			}
		} else {
			fmt.Printf("%d", i)
			// Menambahkan dan mengatur spasi di tengah
			for j := 1; j <= input; j++ {
				fmt.Print(" ")
			}
			fmt.Printf("%d", i)
		}

		fmt.Println() // Pindah ke baris baru setelah satu baris pola selesai dicetak
	}
}
