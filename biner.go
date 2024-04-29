package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	//x := big.NewInt(123)
	//z := fmt.Sprintf("%b", x)
	//fmt.Println("%b", x)
	n := int64(x)
	fmt.Printf("%b", n)
}
