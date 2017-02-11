package main

import "fmt"

func main() {
	n1 := 17
	fmt.Printf("n1=%T", n1)
	fmt.Println()

	n2 := uint(17)
	fmt.Printf("n2=%T", n2)
	fmt.Println()

	n := 256
	b := byte(n)
	fmt.Printf("b=%v:%T", b, b)
	fmt.Println()

	n = -1
	b = byte(n)
	fmt.Printf("b=%v:%T", b, b)
	fmt.Println()

	b1 := byte(255)
	b = b1 + byte(255)
	fmt.Printf("b=%v:%T", b, b)
	fmt.Println()
}

