package main

import (
	"fmt"
	"../util"
)

func main() {
	util.PrintSeparator(1)
	var a [10]int
	fmt.Printf("a=%v:%t\n", a, a)

	s := make([]int, 10)
	fmt.Printf("s=%v:%t\n", s, s)

	util.PrintSeparator(2)

	fmt.Printf("len(s): %v\n", len(s))
	fmt.Printf("cap(s): %v\n", cap(s))

	s2 := make([]int, 5, 10)
	fmt.Printf("len(s2): %v\n", len(s2))
	fmt.Printf("cap(s2): %v\n", cap(s2))

	s3 := a[0:2]
	fmt.Printf("len(s3): %v\n", len(s3))
	fmt.Printf("cap(s3): %v\n", cap(s3))

	util.PrintSeparator(3)

	fmt.Println("あいうえお"[3:9])

	util.PrintSeparator(4)

	s4 := []int{1, 2, 3}
	s4 = append(s4, 4)
	fmt.Println(s4)
	s4 = append(s4, 5, 6, 7)
	fmt.Println(s4)
	s5 := []int{8, 9, 10}
	s4 = append(s4, s5...)
	fmt.Println(s4)

	var b []byte
	b = append(b, "あいうえお"...)
	b = append(b, "かきくけこ"...)
	b = append(b, "さしすせそ"...)
	fmt.Println(b)

	util.PrintSeparator(5)

	s = make([]int, 0, 0)
	fmt.Printf("(A) len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("(B) len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, []int{2, 3, 4}...)
	fmt.Printf("(C) len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, 5)
	fmt.Printf("(D) len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, 6, 7, 8, 9)
	fmt.Printf("(E) len=%d, cap=%d\n", len(s), cap(s))

	s = make([]int, 1024, 1024)
	fmt.Printf("(A') len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, 0)
	fmt.Printf("(B') len=%d, cap=%d\n", len(s), cap(s))

	util.PrintSeparator(6)

	s1 := []int{1, 2, 3, 4, 5}
	s2 = []int{10, 11}
	n := copy(s1, s2)
	fmt.Printf("n=%v, s1=%v\n", n, s1)

	s1 = []int{1, 2, 3, 4, 5}
	s2 = []int{10, 11, 12, 13, 14, 15, 16}
	n = copy(s1, s2)
	fmt.Printf("n=%v, s1=%v\n", n, s1)

	b = make([]byte, 9)
	n = copy(b, "あいうえお")
	fmt.Println(n, b)

	util.PrintSeparator(7)

	a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 = a[2:4]
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v\n", s1, len(s1), cap(s1))

	s2 = a[2:4:4]
	fmt.Printf("s2=%v, len(s2)=%v, cap(s2)=%v\n", s2, len(s2), cap(s2))

	s3 = a[2:4:6]
	fmt.Printf("s3=%v, len(s3)=%v, cap(s3)=%v\n", s3, len(s3), cap(s3))

	util.PrintSeparator(8)

	str := []string{"Apple", "Banana", "Cherry"}
	for i, v := range str {
		fmt.Printf("[%d] => %s\n", i, v)
	}

	util.PrintSeparator(9)

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum())
	fmt.Println(sum([]int{1, 2, 3}...))

	util.PrintSeparator(10)

	aPowAsValue := [3]int{1, 2, 3}
	powAsValue(aPowAsValue)
	fmt.Println(aPowAsValue)

	aPow := []int{1, 2, 3}
	pow(aPow)
	fmt.Println(aPow)

	var (
		array [3]int
		slice []int
	)
	fmt.Println(array)
	fmt.Println(slice == nil)
}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func powAsValue(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func pow(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}
