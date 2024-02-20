package main

import "fmt"

func main() {
	// Method 1: Declare and initialize variables separately
	var a int
	a = 10

	// Method 2: Declare and initialize variables in a single line
	var b int = 20

	// Method 3: Type inference
	var c = 30

	// Method 4: Multiple variable declaration
	var d, e int = 40, 50

	// Method 5: Short variable declaration
	f := 60

	// Method 6: Multiple variable declaration with initial values
	var (
		name = "Abhay"
		age  = 22
	)

	// Method 7: Multiple variable declaration
	var (
		degree     string
		passedYear int
	)

	degree = "B.Tech"
	passedYear = 2024

	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)
	fmt.Println("Value of c:", c)
	fmt.Println("Value of d:", d)
	fmt.Println("Value of e:", e)
	fmt.Println("Value of f:", f)

	fmt.Println("\nName:", name, "| Age:", age)
	fmt.Println("Degree:", degree, "| Passed Year:", passedYear)
}
