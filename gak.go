package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&n)
    fmt.Printf("Faktor dari %d adalah: ", n)
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            fmt.Printf("%d ", i)
        }
    }
}