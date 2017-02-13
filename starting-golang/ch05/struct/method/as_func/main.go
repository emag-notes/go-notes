package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) ToString() string {
	return fmt.Sprintf("[%d,%d]", p.X, p.Y)
}

func main() {
	f := (*Point).ToString
	fmt.Printf("%v:%T\n", f, f)
	fmt.Println(f(&Point{X: 7, Y: 11}))

	fmt.Println((*Point).ToString(&Point{X:11, Y:33}))
}
