package main

import "fmt"

type MyError struct {
	Message string
	ErrorCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrorCode: 1234}
}

///////////////////////////

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

func Println(s Stringify) {
	fmt.Println(s.ToString())
}

///////////////////////////

type T struct {
	Id int
	Name string
}

func (t *T) String() string {
	return fmt.Sprintf("<<%d, %s>>", t.Id, t.Name)
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrorCode)
	}

	vs := []Stringify {
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "XXX-0123", Model: "PX512"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
		Println(v)
	}

	t := &T{Id: 10, Name: "Taro"}
	fmt.Println(t)
}
