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
	//fmt.Println("CPU:\t\t", runtime.NumCPU())
	//fmt.Println("Goroutines:\t\t", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo: ", i)
	}

	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar: ", i)
	}
	wg.Done()
}
