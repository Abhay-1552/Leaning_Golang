package main

import "fmt"

func main() {
    // Standard for loop
    fmt.Println("Standard for loop:")
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // For loop as a while loop
    fmt.Println("\nFor loop as a while loop:")
    sum := 0
    for sum < 100 {
        sum += 10
    }
    fmt.Println("Sum:", sum)

    // For loop with range
    fmt.Println("\nFor loop with range:")
    numbers := []int{1, 2, 3, 4, 5}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // For loop without a condition (infinite loop with break)
    fmt.Println("\nFor loop without a condition (infinite loop with break):")
    count := 0
    for {
        count++
        fmt.Println(count)
        if count == 5 {
            break
        }
    }
}
