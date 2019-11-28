package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//fmt.Println("OS:\t\t", runtime.GOOS)
	//fmt.Println("Architecture:\t\t", runtime.GOARCH)
	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t\t", runtime.NumGoroutine())
	var counter = 0
	const gs = 100
	var mu sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:\t\t", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Goroutines:\t\t", runtime.NumGoroutine())
	fmt.Println("Counter:\t\t", counter)

}
