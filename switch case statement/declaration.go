package main

import "fmt"

func main() {
	// Example 1: Basic switch case
	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's a banana")
	case "orange":
		fmt.Println("It's an orange")
	default:
		fmt.Println("Unknown fruit")
	}

	// Example 2: switch case with fallthrough
	// The fallthrough statement then transfers control to the next case block without considering the condition.
	// In this case, it jumps to case 3.
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	}

	// Example 3: switch case with expression
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	case score >= 60:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
}
