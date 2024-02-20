package main

import "fmt"

func main() {
    // Example 1: Basic if-else
    num := 10
    if num > 0 {
        fmt.Println("Number is positive")
    } else {
        fmt.Println("Number is non-positive")
    }

    // Example 2: if-else if ladder
    score := 75
    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    } else if score >= 60 {
        fmt.Println("Grade: D")
    } else {
        fmt.Println("Grade: F")
    }

    // Example 3: Short statement before condition
    if x := 5; x < 10 {
        fmt.Println("x is less than 10")
    }

    // Example 4: Nested if-else
    age := 25
    if age >= 18 {
        if age < 60 {
            fmt.Println("Adult")
        } else {
            fmt.Println("Senior Citizen")
        }
    } else {
        fmt.Println("Minor")
    }
}
