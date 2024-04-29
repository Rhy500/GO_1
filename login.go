package main

import "fmt"

fun main() {
	var username, password string 
	var jumlah int
	jumlah = 0

	for jumlah!{
		jumlah = jumlah + 1
		fmt.Scan(&username, &password)
		jumlah = username =="admin" && password == "admin"
	}
	fmt.Println(jumlah - 1)
	fmt.Println("login berhasil")
}