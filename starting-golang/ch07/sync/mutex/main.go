package main

import (
	"time"
	"fmt"
	"sync"
)

var st struct {
	A, B, C int
}

var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	mutex.Lock()

	st.A = n
	sleep()
	st.B = n
	sleep()
	st.C = n
	sleep()

	fmt.Println(st.A, st.B, st.C)

	mutex.Unlock()
}

func sleep() {
	time.Sleep(time.Microsecond)
}

func main() {
	mutex = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		} ()
	}
	for {
	}
}