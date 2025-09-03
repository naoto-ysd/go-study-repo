package main

import "fmt"

// 変数と定数の学習
func main() {
	// 変数の宣言方法
	var name string = "太郎"
	var age int = 25

	// 短い宣言構文
	city := "東京"
	isStudent := true

	// 複数変数の宣言
	var (
		firstName string = "山田"
		lastName  string = "太郎"
	)

	// 定数の宣言
	const pi = 3.14159
	const greeting = "こんにちは"

	// ポインタの宣言
	var ptr *string = &name

	// スライス
	var slice = []int{1, 2, 3, 4, 5}

	// 出力
	fmt.Printf("名前: %s\n", name)
	fmt.Printf("年齢: %d\n", age)
	fmt.Printf("都市: %s\n", city)
	fmt.Printf("学生: %t\n", isStudent)
	fmt.Printf("フルネーム: %s %s\n", firstName, lastName)
	fmt.Printf("円周率: %.2f\n", pi)
	fmt.Println(greeting)
	fmt.Printf("名前のポインタ: %p, ポインタの指す値: %s\n", ptr, *ptr)
	fmt.Printf("スライスの値: %v\n", slice)

	// sliceの回数だけ繰り返す
	for i := 0; i < len(slice); i++ {
		fmt.Printf("スライスの要素 %d: %d\n", i, slice[i])
	}

	// range
	for i, v := range slice {
		fmt.Printf("スライスの要素 %d: %d\n", i, v)
	}
}
