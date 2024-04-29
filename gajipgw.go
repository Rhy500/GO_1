package main

import "fmt"

func main() {
	var jabatan string
	var masakrj, jumank int
	fmt.Scan(&jabatan, &masakrj, &jumank)

	if jabatan == "staf" {
		if masakrj < 5 {
			fmt.Println(4000 + 0)
		} else if masakrj > 10 {
			fmt.Println(5000 + jumank*100)
		} else {
			fmt.Println(4000 + jumank*100)
		}

	} else if jabatan == "Manager" {
		if masakrj > 10 {
			fmt.Println(10000 + jumank*300)
		} else {
			fmt.Println(8500 + jumank*300)
		}
	} else {
		fmt.Println(20000 + jumank*500)
	}
}
