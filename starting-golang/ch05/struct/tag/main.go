package main

import (
	"reflect"
	"fmt"
	"encoding/json"
)

type User struct {
	Id int `json:"user_id"`
	Name string `json:"user_name"`
	Age uint `json:"user_age"`
}

func main() {
	u := User{Id: 1, Name: "Taro", Age: 32}

	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
	}

	bs, _ := json.Marshal(u)
	fmt.Println(string(bs))

}
