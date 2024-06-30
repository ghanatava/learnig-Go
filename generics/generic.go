package main

import "fmt"

func main() {
	fmt.Printf("Has a %v \n", Has([]string{"a", "b"}, "a"))
	fmt.Printf("Has 1 %v \n", Has([]int32{1, 2, 3}, 2))

	fmt.Printf("Mylist %v\n", NewEmptyList[int]()) //[int] is used to explicitly tell the type else compiler can't infer it.

	PrintThings("abs", "def", 123.45, 12)
}

func NewEmptyList[T any]() []T {
	return make([]T, 0)
}

func Has[T comparable](list []T, value T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func PrintThings[A, B any, C ~int](a1, a2 A, b B, c C) {
	fmt.Printf("%v %v %v %v\n", a1, a2, b, c)
}
