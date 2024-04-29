package main

import "fmt"

func main() {
	var b1, b2, b3, b4, rata int64
	//var rata float
	fmt.Scan(&b1, &b2, &b3, &b4)

	rata = (b1 + b2 + b3 + b4) / 4
	if rata >= int(3.50) {
		fmt.Println("bagus")
	} else if int(rata <= 1.50) {
		fmt.Println("sedang")
	} else {
		fmt.Println("kurang")
	}
}
