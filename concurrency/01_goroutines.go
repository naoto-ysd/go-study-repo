package main

import (
    "fmt"
    "time"
)

// ゴルーチンで実行される関数
func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Printf("Hello from %s (%d)\n", name, i+1)
        time.Sleep(100 * time.Millisecond)
    }
}

// チャネルを使った例
func sendNumbers(ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- i
        fmt.Printf("送信: %d\n", i)
        time.Sleep(100 * time.Millisecond)
    }
    close(ch)
}

func main() {
    fmt.Println("=== ゴルーチンの例 ===")
    
    // ゴルーチンの開始
    go sayHello("ゴルーチン1")
    go sayHello("ゴルーチン2")
    
    // メインゴルーチンも同時に動作
    sayHello("メイン")
    
    // 少し待機してゴルーチンが完了するのを待つ
    time.Sleep(500 * time.Millisecond)
    
    fmt.Println("\n=== チャネルの例 ===")
    
    // チャネルの作成
    ch := make(chan int)
    
    // ゴルーチンでデータを送信
    go sendNumbers(ch)
    
    // チャネルからデータを受信
    for number := range ch {
        fmt.Printf("受信: %d\n", number)
    }
    
    fmt.Println("プログラム終了")
}
