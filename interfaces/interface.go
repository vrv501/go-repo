package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Triangle struct {
	length float64
	height float64
}

func (t *Triangle) Area() float64 {
	return t.length * t.height
}

var _ Shape = (*Triangle)(nil)

type Square struct {
	side float64
}

func (s *Square) Area() float64 {
	return math.Pow(s.side, 2)
}

var _ Shape = (*Square)(nil)

// Notice below any "type" can implement an interface
type MyOwnDataType int32

func (md *MyOwnDataType) Area() float64 {
	return float64(0)
}

func isExists(list []interface{}, val interface{}) bool {
	for _, ele := range list {
		if val == ele {
			return true
		}
	}
	return false
}

func main() {
	var s Shape
	// Interface is nothing but tuple of (value, concrete-type)
	// type can be anything including go built-in-types
	// Can also be your own structs
	s = &Triangle{length: 2, height: 2}
	fmt.Println(s.Area())

	s = &Square{side: 2}
	fmt.Println(s.Area())

	// For a type to implicitly implemet
	// an interface means it implements all the methods
	// presnet in interface with same signatures

	// The below statement is type assertion
	// If interface is of type (*Square), then
	// assertion would return underlying value &Square{side: 2
	shapeType, ok := s.(*Square)
	if ok {
		fmt.Printf("%T %+v\n", shapeType, shapeType)
	}

	// A interface without methods is empty interface
	// Technically any type implements empty interface
	// Thus interface{} can denote any type - (val, anyConcrete-Type)
	var anyType interface{} = 3.456
	fmt.Println(anyType)

	// Note comparision operators can only be used with same dataTypes
	fmt.Println(isExists([]interface{}{1, 2, 3.56, 3.456}, anyType))

	// Go requires explicit type-casting and doesn't do
	// implicit type-casting
	fmt.Println(isExists([]interface{}{1, 2, 3.56, float32(3.456)}, anyType))
}
