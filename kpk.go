package main

import "fmt"

func main() {
	var n, m, i, j, a, b int
	fmt.Scan(&n, &m)
	// cara 1
	//fmt.Println(N * M)

	// cara ke 2
	j = n * m
	for i = 1; i <= j; i++ {
		a = i % n
		b = i % m
		if a == 0 && b == 0 {
			j = 0
		}
	}
	fmt.Println(i - 1 )
}
