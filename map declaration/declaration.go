package main

import "fmt"

func main() {
	// 1. Using make() function
	var mapUsingMake = make(map[string]int)

	// 2. Using map literals
	var mapUsingLiteral = map[string]int{"a": 1, "b": 2, "c": 3}

	// 3. Short variable declaration
	mapShortDeclaration := make(map[string]int)

	// 4. Nil map
	var nilMap map[string]int

	// 5. Type alias
	type StringIntMap map[string]int
	var mapTypeAlias StringIntMap
	mapTypeAlias = make(StringIntMap)

	// Adding key-value pairs to maps
	mapUsingMake["x"] = 10
	mapUsingLiteral["y"] = 20
	mapShortDeclaration["z"] = 30
	nilMap = make(map[string]int)
	nilMap["w"] = 40
	mapTypeAlias["v"] = 50

	// Printing maps
	fmt.Println("mapUsingMake:", mapUsingMake)
	fmt.Println("mapUsingLiteral:", mapUsingLiteral)
	fmt.Println("mapShortDeclaration:", mapShortDeclaration)
	fmt.Println("nilMap:", nilMap)
	fmt.Println("mapTypeAlias:", mapTypeAlias)
}
