package main

import "fmt"

type MyInt int

func (m MyInt) Plus(i int) int {
	return int(m) + i
}

type IntPair [2]int

func (ip IntPair) First() int {
	return ip[0]
}

func (ip IntPair) Last() int {
	return ip[1]
}

type Strings []string
func (s Strings) Join(d string) string {
	sum := ""
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(MyInt(4).Plus(2))

	ip := IntPair{1, 2}
	fmt.Println(ip.First(), ip.Last())

	fmt.Println(Strings{"A", "B", "C"}.Join(","))

}
