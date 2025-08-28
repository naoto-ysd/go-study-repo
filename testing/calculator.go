package calc

import (
    "errors"
)

// 計算機の関数群
func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

func Multiply(a, b int) int {
    return a * b
}

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("ゼロで割ることはできません")
    }
    return a / b, nil
}
