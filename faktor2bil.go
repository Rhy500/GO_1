package main 

import "fmt"

func main() {
	var x,y int
	var total bool

	total = 0
	for total == 0 {
		total = x % y
		fmt.Println(total)
	}
}