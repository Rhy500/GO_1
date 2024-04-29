package main

import "fmt"

func main() {
	var n, total int
	var m1, m2, m3, m4, m5 int
	fmt.Scan(&n)
	if n == 4 {
		fmt.Scan(&m1, &m2, &m3, &m4)
		total = m1 + m2 + m3 + m4
		fmt.Println(total)
	} else {
		if n == 5 {
			fmt.Scan(&m1, &m2, &m3, &m4, &m5)
			total = m1 + m2 + m3 + m4 + m5
			fmt.Println(total)
		}
	}
}
