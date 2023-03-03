
package main

import (
    "fmt"
    "time"
)

func main() {
    var birthYear int
    fmt.Print("Enter your birth year: ")
    fmt.Scan(&birthYear)

    currentYear := time.Now().Year()
    age := currentYear - birthYear

    fmt.Println("Your age is:", age)
}

// Scenario 
//Age Calculator: Write a program that prompts the user for their birth year 
// and calculates their age using variables to store the values.
