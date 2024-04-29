package main

import "fmt"

func main() {
	var total int64
	for i := 2; i <= 50; i++ {
		total += 1
	}
	fmt.Println(total)
}
