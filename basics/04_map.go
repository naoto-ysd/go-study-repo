package main

import "fmt"

// 変数と定数の学習
func main() {
	totalWins := map[string]int{}

	fmt.Println(totalWins == nil)
	fmt.Println(totalWins["abc"])
	totalWins["abc"] = 3
	fmt.Println(totalWins["abc"])

	mayBeZero := map[string]int{}
	mayBeZero["tom"] = 100
	fmt.Printf("tomは100: %d\n", mayBeZero["tom"])
	fmt.Printf("Bob is 0: %d\n", mayBeZero["bob"])

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	delete(m, "world")
	fmt.Printf("mの値: %d\n", m["world"])

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
}
