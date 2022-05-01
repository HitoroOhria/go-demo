package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	fmt.Printf("main...\n")

	meeting()
	channel()
}
