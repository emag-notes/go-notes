package main

import "fmt"

func main() {
	var n int = 5
	fmt.Printf("n=%d:%T\n", n, n)

	var (
		x, y int
		name string = "String?"
	)
	x, y = 1, 2
	name = "String!"

	fmt.Printf("x=%v , y=%v, name=%v\n", x, y, name)

	i := 1
	fmt.Printf("i=%d:%T\n", i, i)

	one := one()
	fmt.Printf("one=%d:%T\n", one, one)
}

func one() int {
	return 1
}