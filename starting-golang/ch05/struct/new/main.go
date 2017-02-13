package main

import "fmt"

type Person struct {
	Id int
	Name string
	Area string
}

func main() {
	p := new(Person)
	p.Id = 10
	p.Name = "Taro"
	p.Area = "Japan"
	fmt.Println(p)
}
