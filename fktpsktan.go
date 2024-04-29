package main

import "fmt"

func main() {
	var n, m, i, j int

	fmt.Scan(&n, &m)

	for i = 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println(" ")
	for j = 1; j <= m; j++ {
		if m%j == 0 {
			fmt.Print(j, " ")
		}
	}
	fmt.Println("Abil nilai yg sama dari pengfaktoran number 1 dan number 2 ")
}
