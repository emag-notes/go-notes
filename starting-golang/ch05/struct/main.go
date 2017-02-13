package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Feed struct {
	Name string
	Amount uint
}

type Animal struct {
	Name string
	Feed
}

func main() {
	fmt.Println("------------------------")
	printPoint()

	fmt.Println("------------------------")
	printAnimal()
}

func printPoint() {
	pt := Point{}
	fmt.Println(pt)

	pt.X = 10
	pt.Y = 8
	fmt.Println(pt)

	pt = Point{1, 2}
	fmt.Println(pt)

	pt = Point{X: 1, Y: 2}
	fmt.Println(pt)

	pt = Point{Y: 2}
	fmt.Println(pt)
}

func printAnimal() {
	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name: "Banana",
			Amount: 10,
		},
	}
	fmt.Println(a)
	fmt.Println(a.Name)
	fmt.Println(a.Feed)
	fmt.Println(a.Feed.Name)
	a.Feed.Amount = 20
	//a.Amount = 100
	fmt.Println(a.Feed.Amount)
}