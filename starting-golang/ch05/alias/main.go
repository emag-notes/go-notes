package main

import "fmt"

type Ints []int
type Callback func(i int) int

func Sum(ints Ints, callback Callback) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return callback(sum)
}

func main() {
	n := Sum(
		Ints{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)
	fmt.Println(n)
}
