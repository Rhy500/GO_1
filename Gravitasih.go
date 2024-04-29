package main

import "fmt"

func main() {
	var merkurius, venus, bumi, mars, yupiter, saturnus, uranus, neptunus float64
	var g_merkurius, g_venus, g_mars, g_yupiter, g_saturnus, g_uranus, g_neptunus, g_bumi float64
	var berat float64
	fmt.Scan(&berat)
	g_bumi = 9.8
	g_merkurius = 0.38
	g_venus = 0.91
	g_mars = 0.38
	g_yupiter = 2.37
	g_saturnus = 0.92
	g_uranus = 0.89
	g_neptunus = 1.13

	bumi = berat * g_bumi
	merkurius = bumi * g_merkurius
	venus = bumi * g_venus
	mars = bumi * g_mars
	yupiter = bumi * g_yupiter
	saturnus = bumi * g_saturnus
	uranus = bumi * g_uranus
	neptunus = bumi * g_neptunus
	//fmt.Println(int64 (merkurius, venus, bumi, mars, yupiter, saturnus, uranus, neptunus))
	fmt.Println(int(merkurius))
	fmt.Println(int(venus))
	fmt.Println(int(bumi))
	fmt.Println(int(mars))
	fmt.Println(int(yupiter))
	fmt.Println(int(saturnus))
	fmt.Println(int(uranus))
	fmt.Println(int(neptunus))
}
