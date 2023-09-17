package main

import "fmt"

// Generic type parameters
func isExists[T comparable](list []T, val T) bool {
	for _, ele := range list {
		if val == ele {
			return true
		}
	}

	return false
}

// Generics can also be used with owntypes
type LinkedList[T any] struct {
	value T
}

func main() {
	fmt.Println(isExists([]int{1, 2, 3}, 3))

	// Both will be sent as same dataType
	fmt.Println(isExists([]float32{1, 2, 3}, 3))

	// Below returns errors during compileTime
	// Which is why generics are way to go for anyType
	// rather than using interfaces
	// fmt.Println(isExists([]float32{1, 2, 3}, int(3)))

	// comparable allows any dataTypes as longas they allow comparision
	// using == and !=

	// any is just alias to interface{}
}
