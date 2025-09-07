package main

import "fmt"

// 変数と定数の学習
func main() {
	totalWins := map[string]int{}

	fmt.Println(totalWins == nil)
	fmt.Println(totalWins["abc"])
	totalWins["abc"] = 3
	fmt.Println(totalWins["abc"])
}
