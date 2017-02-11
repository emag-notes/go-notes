package main

import (
	"fmt"
	"math"
)

func main() {
	ui_1 := uint32 (400000000)
	ui_2 := uint32(4000000000)
	ui_3 := uint32(4200000000)

	sum := ui_1 + ui_2
	fmt.Printf("%d + %d = %d", ui_1, ui_2, sum)
	fmt.Println()
	fmt.Printf("%d < %d = %v", ui_3, sum, ui_3 < sum)
	fmt.Println()

	fmt.Printf("uint32 max value = %d:%T", math.MaxUint32, math.MaxUint32)
	fmt.Println()
}