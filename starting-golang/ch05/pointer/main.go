package main

import "fmt"

func main() {
	var i int = 5
	p := &i
	fmt.Printf("type=%T, address=%p, value=%v\n", p, p, *p)


	*p = 10
	fmt.Println(i)

	pp := &p
	fmt.Printf("%T\n", pp)
}
