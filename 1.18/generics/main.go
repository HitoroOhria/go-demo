package main

import "fmt"

func main() {
	compare()
}

func compare() {
	// true
	fmt.Println(include([]string{"foo", "bar"}, "bar"))

	// false
	fmt.Println(include([]string{"foo", "bar"}, "baz"))

	// values := []*ValueObject{
	// 	&ValueObject{Value: "foo"},
	// 	&ValueObject{Value: "bar"},
	// 	&ValueObject{Value: "baz"},
	// }
	//
	// fmt.Println(values, &ValueObject{Value: "foo"})
}

func include[T comparable](list []T, value T) bool {
	for _, element := range list {
		if element == value {
			return true
		}
	}

	return false
}

type ValueObject[T comparable] struct {
	Value T
}

// WIP interface is (or embeds) comparable
func includeForValueObj[T ValueObject[comparable]](list []T, value T) bool {
	for _, element := range list {
		// WIP undefined (type T has no field or method Value)[comparable]
		if element.Value == value.Value {
			return true
		}
	}

	return false
}
