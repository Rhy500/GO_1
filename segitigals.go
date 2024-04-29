package main

import "fmt"

func main() {
	var s1, s2, s3 int
	fmt.Scan(&s1, &s2, &s3)

	if s1 == s2 && s2 == s3 && s3 == s2 {
		fmt.Println("segitiga sama sisi")
	} else if s1 <= s2 && s2 >= s3 && s3 == s1 {
		fmt.Println("segitiga sama kaki")
	} else {
		fmt.Println("segitiga sembarang")
	}
}
