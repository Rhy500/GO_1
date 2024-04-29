package main

import (
	"fmt"
)

func main() {
	var s1, s2, s3, s4, s5 int
	fmt.Scan(&s1, &s2)

	s3 = s2 + s1
	s4 = s3 + s2
	s5 = s4 + s3
	fmt.Println(s3, s4, s5)
}
