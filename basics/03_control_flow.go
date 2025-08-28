package main

import "fmt"

// 制御構文の学習
func main() {
    // if文
    age := 20
    if age >= 20 {
        fmt.Println("成人です")
    } else if age >= 13 {
        fmt.Println("ティーンエイジャーです")
    } else {
        fmt.Println("子供です")
    }
    
    // for文（基本形）
    fmt.Println("\n1から5まで:")
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()
    
    // for文（while風）
    fmt.Println("\nカウントダウン:")
    count := 3
    for count > 0 {
        fmt.Printf("%d ", count)
        count--
    }
    fmt.Println("発射!")
    
    // for range（スライス）
    fruits := []string{"りんご", "バナナ", "オレンジ"}
    fmt.Println("\n果物:")
    for index, fruit := range fruits {
        fmt.Printf("%d: %s\n", index, fruit)
    }
    
    // switch文
    day := "月曜日"
    switch day {
    case "月曜日":
        fmt.Println("週の始まりです")
    case "金曜日":
        fmt.Println("週末が近いです")
    case "土曜日", "日曜日":
        fmt.Println("週末です")
    default:
        fmt.Println("平日です")
    }
}
