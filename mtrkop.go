package main

import "fmt"

func main() {
	var posig int
	var stskp, stsg bool
	fmt.Scan(&posig, &stskp, &stsg)
	if posig == 0 || stskp == true && stsg == false {
		fmt.Println("Mesin menyala dan motor tidak berjalan")
	} else if posig == 1 || posig == 2 || posig == 3 || posig == 4 && stskp == false && stsg == true {
		fmt.Println("Motor berjalan")
	} else {
		fmt.Println("Mesin mati")
	}
}
