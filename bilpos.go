package main

import (
	"fmt"
)

func main() {
	var sum int
	for {
		var num int
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		sum += num
	}
	fmt.Println(sum)
}
