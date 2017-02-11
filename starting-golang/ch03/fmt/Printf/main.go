package main

import (
	"fmt"
)

func main() {
	fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17)
	fmt.Printf("%d年%d月%d日\n", 2017, 2)
	fmt.Printf("%d年%d月%d日\n", 2017, 2, 11, 12)

	fmt.Println("\n------------------------------")

	fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3})
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3})
}
