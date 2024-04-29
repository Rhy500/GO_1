package main 

import "fmt"

func main() {
	var s1, s2, s3 int
	var hasil bool
	fmt.Scan(&s1, &s2, &s3)
	hasil=((s1+s2 > s3) && (s2+s3 > s1) && (s1+s3 > s2))
	fmt.Println(hasil)
}