package main

import "fmt"

func main() {
	var h1, m1 int
	// h1 jam , m1 menit pada saat parkir
	var h2, m2 int
	// h2 jam dan m2 menit selesai parkir
	fmt.Scan(&h1, &m1, &h2, &m2)
	if h2 < h1 {
		fmt.Println((12-h1)+h2, "jam", (m2 - m1), "menit")
	} else {
		fmt.Println((h2 - h1), "jam", (m2 - m1), "menit")
	}
}
