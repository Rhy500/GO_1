package main

import "fmt"

func main() {
	var n int
	var item1, item2, item3, item4, kesiapantim bool
	fmt.Scan(&n)
	kesiapantim = true

	for i := 1; i < n; i++ {
		fmt.Scan(&item1, &item2, &item3, &item4)
		kesiapantim = item1 && item2 && item3 && item4
	}
	fmt.Println(kesiapantim)
}
