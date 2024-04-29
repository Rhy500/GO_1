package main

import "fmt"

func main() {
	var n, m, x, y, cangkir int
	fmt.Scan(&n, &m, &x, &y)
	cangkir = 0
	for x <= n && y <= m {
		n = n - x
		m = m - y
		cangkir = cangkir + 1
	}
	fmt.Println(cangkir)
}
