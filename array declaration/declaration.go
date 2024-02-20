package main

import "fmt"

func main() {
	// Standard Array Declaration - default value is 0
	var numbers1 [5]int

	// Array Declaration with Initialization
	var numbers2 = [5]string{"a", "b", "c", "d", "e"}

	// Array Declaration with Implicit Sizing
	numbers3 := [...]int{1, 2, 3, 4, 5}

	// Array Declaration with Specific Index Initialization
	numbers4 := [5]int{0: 10, 2: 20, 4: 30}

	// Multidimensional Array Declaration
	var matrix [3][3]int

	// Array Declaration with Specific Type
	type MyArray [5]int
	var numbers5 MyArray

	// Printing the arrays
	fmt.Println("numbers1:", numbers1)
	fmt.Println("numbers2:", numbers2)
	fmt.Println("numbers3:", numbers3)
	fmt.Println("numbers4:", numbers4)
	fmt.Println("matrix:", matrix)
	fmt.Println("numbers5:", numbers5)
}
