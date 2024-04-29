package main

import "fmt"

func main() {
	var k int64
	fmt.Scanf("%c", &k)

	if k >= '9' || k >= 'E' || k >= 't' {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
