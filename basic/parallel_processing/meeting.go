package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func meeting() {
	fmt.Printf("start meeting...\n")

	manualMeeting()
	waitGroupMeeting()
}

func printNum() {
	time.Sleep(1 * time.Second)
	fmt.Printf("the number is %d\n", rand.Intn(10))
}

func manualMeeting() {
	fmt.Printf("start manula Meeting\n")
	go printNum()

	// This line is required.
	// Because main goroutine ends without waiting for the end of sub goroutine.
	time.Sleep(2 * time.Second)
}

func waitGroupMeeting() {
	fmt.Printf("start WaiteGroup Meeting\n")

	var wg sync.WaitGroup
	// Add counter of WaiteGroup.
	wg.Add(1)

	go func() {
		// Minus counter of WaiteGroup.
		defer wg.Done()
		printNum()
	}()

	// Block main goroutine and wait until counter of WaitGroup reaches 0.
	wg.Wait()
}
