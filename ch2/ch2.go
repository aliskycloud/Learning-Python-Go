package main

import "fmt"

func main() {
    var celsius float64
    fmt.Print("Enter temperature in Celsius: ")
    fmt.Scan(&celsius)

    fahrenheit := (celsius * 9 / 5) + 32

    fmt.Println(celsius, "Celsius is equal to", fahrenheit, "Fahrenheit.")
}
