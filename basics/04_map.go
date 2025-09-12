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
}
