package main

import (
	"fmt"
)

func main() {
	var input string
	var stop bool
	stop = false

	for !stop {
		fmt.Scan(&input)
		stop = input == "ada"
		if !stop {
			fmt.Println("cari beasiswa")
		}
	}
	fmt.Println("pencarian selesai")
}
