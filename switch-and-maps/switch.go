package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Initialisation + selection
	// no breaks required
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println(os)
	}

	// The case statements need not be constants
	// and can be constants
	randValue := 2
	targetValue := 4
	switch {
	case randValue+1 == targetValue:
		fmt.Println("Equal")
	case randValue+2 == targetValue:
		fmt.Println("Equal")
	default:
		fmt.Println("Not equal")
	}

	// You can use type switch
	var x interface{} = 6.789
	switch x.(type) {
	case int32:
		fmt.Println("integer")
	case float64:
		fmt.Println("decimal")
	default:
		fmt.Printf("Unknown")
	}

	// You can check if key exists in map
	m := map[string]int{}
	val, ok := m["key"]
	if ok {
		fmt.Println(val)
	}

}
