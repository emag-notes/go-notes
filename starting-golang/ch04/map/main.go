package main

import (
	"fmt"
	"../util"
)

func main() {
	util.PrintSeparator(1)

	m1 := make(map[int]string)
	m1[1] = "US"
	m1[81] = "Japan"
	m1[86] = "China"

	fmt.Println(m1)

	util.PrintSeparator(2)

	m2 := make(map[string]string)
	m2["Yamada"] = "Taro"
	m2["Sato"] = "Hanako"
	m2["Yamada"] = "Jiro"

	fmt.Println(m2)

	util.PrintSeparator(3)

	m3 := make(map[float64]int)
	m3[1.0000000000000000000000001] = 1
	m3[1.0000000000000000000000002] = 2
	m3[1.0000000000000000000000003] = 3

	fmt.Println(m3)

	util.PrintSeparator(4)

	m4 := map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	fmt.Println(m4)

	m4 = map[int]string{
		1: "Taro",
		2: "Hanako",
		3: "Jiro",
	}
	fmt.Println(m4)

	util.PrintSeparator(5)

	m5 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}
	fmt.Println(m5)

	m5 = map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	fmt.Println(m5)

	util.PrintSeparator(6)

	m6 := map[int]map[float64]string{
		1: {3.14: "円周率"},
	}
	fmt.Println(m6)

	util.PrintSeparator(7)

	m7 := map[int]string{1: "A", 2: "B", 3: "C"}
	fmt.Println(m7[1])
	fmt.Println(m7[9])

	util.PrintSeparator(8)

	m8 := map[int]string{1: "A", 2: "B", 3: "C"}

	s8_1, ok8_1 := m8[1]
	fmt.Println(s8_1, ok8_1)
	s8_2, ok8_2 := m8[9]
	fmt.Println(s8_2, ok8_2)
	_, ok8_3 := m8[3]
	fmt.Println(ok8_3)

	if _, ok := m8[2]; ok {
		fmt.Println("m8[2]", m8[2])
	}

	util.PrintSeparator(9)

	m9 := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}

	for k, v := range m9 {
		fmt.Printf("%d => %s\n", k, v)
	}

	util.PrintSeparator(10)

	m10 := map[int]string{1: "A", 2: "B", 3: "C"}
	fmt.Println(len(m10))
	m10[4] = "D"
	m10[5] = "E"
	fmt.Println(len(m10))

	util.PrintSeparator(11)

	m11 := map[int]string{1: "A", 2: "B", 3: "C"}

	delete(m11, 2)

	fmt.Println(m11)
}
