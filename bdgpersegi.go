package main

import "fmt"

func main() {
	var n, luas, keliling, sisi, lebar int
	fmt.Scan(&n)
	luas = 0
	keliling = 0
	for i := n; i >= n; i++ {
		fmt.Scan(&sisi)
		fmt.Scan(&lebar)
		luas = sisi * lebar
		keliling = 2 * (sisi + lebar)
		fmt.Println("%d %d\n", luas, keliling)
	}

}
