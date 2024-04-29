package main

import "fmt"

func main() {
	var jam, menit int
	var jarak float64
	fmt.Scan(&jam, &menit, &jarak)

	if jam >= 5 && jam <= 22 {
		if jam >= 5 && (jam <= 8 && menit <= 59) || jam >= 16 && (jam <= 18 && menit <= 59) || jam == 9 || jam == 19 {
			if jarak >= 0 || jarak < 10 {
				fmt.Println(jarak * 5000)
			} else if jarak >= 10 || jarak <= 20 {
				fmt.Println(jarak * 4500)
			}
		} else {
			if jarak > 0 || jarak <= 20 {
				fmt.Println(jarak * 4000)
			}
		}

	} else {
		fmt.Println("maaf kami tidak melayani")
	}
}
