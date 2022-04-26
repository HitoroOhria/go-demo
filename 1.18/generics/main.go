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
}

func include[T comparable](list []T, value T) bool {
    for _, element := range list {
        if element == value {
            return true
        }
    }

    return false
}
