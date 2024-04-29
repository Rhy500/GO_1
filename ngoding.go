package main

import "fmt"

func main() {
	var b, jam, hasilb int
	var rata float64
	fmt.Scan(&b)

	for i := 1; i <= b; i++ {
		fmt.Scan(&jam)
		hasilb = hasilb + jam
	}
	rata = float64(hasilb) / float64(b)
	fmt.Println("%.3f", rata)
}
