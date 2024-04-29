package main

import (
	"fmt"
)

func main() {
	var n, faktorial, i int
	faktorial = 1
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		faktorial = faktorial * i
	}
	fmt.Println(faktorial)
}
