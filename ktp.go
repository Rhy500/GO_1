package main

import "fmt"

func main() {
	var u int
	var n bool
	// u is year old . n is have / not
	fmt.Scan(&u, &n)

	if u >= 17 && n == true {
		fmt.Println("make KTP")
	} else {
		fmt.Println("not make KTP")
	}
}
