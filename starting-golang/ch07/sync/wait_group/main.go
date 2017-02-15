package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine: ", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine: ", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine: ", i)
		}
		wg.Done()
	}()

	wg.Wait()
}
