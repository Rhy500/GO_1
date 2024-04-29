package main

import "fmt"

func main() {
	var membership string
	var a, b, c, d, e int
	var cashback, diskon float64
	fmt.Scan(&membership)
	fmt.Scan(&a, &b, &c, &d, &e)

	if (a%2 != 0) && (b%2 != 0) && (c%2 != 0) && (d%2 != 0) && (e%2 != 0) {
		if membership == "Yes" {
			diskon = (1.7 + (0.15 * 1.7)) * float64((c + d + e))
		} else {
			diskon = 1.7 * float64((c + d + e))
		}
	} else if (a%2 == 0) && (b%2 == 0) && (c%2 == 0) && (d%2 == 0) && (e%2 == 0) {
		if membership == "Yes" {
			cashback = (3.1 + (0.15 * 3.1)) * float64((a + b + c))
		} else {
			cashback = 3.1 * float64((a + b + c))
		}
	} else {
		if membership == "Yes" {
			cashback = (3.1 + (0.15 * 3.1)) * float64((a + b + c))
			diskon = (1.7 + (0.15 * 1.7)) * float64((c + d + e))
		} else {
			cashback = 3.1 * float64((a + b + c))
			diskon = 1.7 * float64((c + d + e))
		}
	}
	if cashback > 35 {
		cashback = 35
	}
	if diskon > 50 {
		diskon = 50
	}
	fmt.Printf("Cashback: %g%% Diskon:%g%% ", cashback, diskon)
}
