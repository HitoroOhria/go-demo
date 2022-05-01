package main

import (
	"fmt"
	"math/rand"
	"time"
)

func channel() {
	fmt.Printf("start channel...\n")

	basic()
	bufferedChannel()
	referenceValue()
	conflictValue()
}

func getNum(c chan<- int) {
	time.Sleep(1 * time.Second)

	num := rand.Intn(10)
	c <- num
}

func basic() {
	fmt.Printf("start basic...\n")

	c := make(chan int)
	defer close(c)

	go getNum(c)
	// Block main goroutine until receive value from channel.
	num := <-c

	fmt.Printf("getting number from channel is %d\n", num)
}

func bufferedChannel() {
	fmt.Printf("start bufferedChannel...\n")

	c := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("finish sub goroutine\n")

		c <- 1
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("finish main gorutine\n")

	// Sync until finishing sub goroutine.
	<-c
}

func referenceValue() {
	fmt.Printf("start reference value...\n")

	// 1. loop processing proceeds (i = 3)
	// 2. executed sub goroutines
	fmt.Printf("bad code\n")
	for i := 0; i < 3; i++ {
		go func() {
			// num is 3
			// num is 3
			// num is 3
			fmt.Printf("num is %d\n", i)
		}()
	}
	time.Sleep(1 * time.Second)

	fmt.Printf("good code\n")
	for i := 0; i < 3; i++ {
		go func(j int) {
			// num is 0
			// num is 2
			// num is 1
			fmt.Printf("num is %d\n", j)
		}(i)
	}
	time.Sleep(1 * time.Second)
}

func conflictValue() {
	fmt.Printf("start reference value...\n")

	fmt.Printf("bad cofe\n")

	src1 := []int{1, 2, 3, 4, 5}
	dst1 := []int{}

	for _, s := range src1 {
		go func(s int) {
			result := s * 2

			// There is time lag between:
			//   1. read from src1
			//   2. write to dst1
			dst1 = append(src1, result)
		}(s)
	}

	time.Sleep(1 * time.Second)
	// [1 2 3 4 5 8]
	// [2 6 8]
	fmt.Printf("dst1 = %v\n", dst1)

	fmt.Printf("good cofe\n")

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := []int{}

	c := make(chan int)

	for _, s := range src2 {
		go func(s int) {
			result := s * 2

			c <- result
		}(s)
	}

	for _ = range src2 {
		dst2 = append(dst2, <-c)
	}

	time.Sleep(1 * time.Second)
	// [10 2 4 6 8]
	// [10 4 6 8 2]
	fmt.Printf("dst2 = %v\n", dst2)
}
