package main

import "fmt"

func ex1() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4, 2, 1}
	var min int = 0

	for i := 0; i < len(l); i++ {
		if i == 0 {
			min = l[i]
		} else if l[i] < min {
			min = l[i]
		}
	}

	fmt.Println("最小値は", min)
}

func ex2() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orage":  80,
		"papaya": 500,
		"kiwi":   90,
	}

	var sum int = 0

	for _, v := range m {
		sum += v
	}

	fmt.Println("合計は", sum)
}

func main() {
	ex1()
	ex2()
}
