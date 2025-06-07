package main

import (
	"fmt"
	"sync"
	"time"
)

func calculate(i int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Millisecond)
	fmt.Println(i * i)
	wg.Done()

}
func Goroutine() {
	var wg sync.WaitGroup
	wg.Add(9)

	for i := 1; i < 10; i++ {
		go calculate(i, &wg)
	}
	wg.Wait()

}
