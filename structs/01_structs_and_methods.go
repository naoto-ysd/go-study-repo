package main

import "fmt"

// 構造体の定義
type Person struct {
    Name string
    Age  int
    City string
}

// メソッドの定義
func (p Person) Introduce() {
    fmt.Printf("こんにちは、私は%sです。%d歳で、%sに住んでいます。\n", 
        p.Name, p.Age, p.City)
}

// ポインタレシーバーを使ったメソッド
func (p *Person) HaveBirthday() {
    p.Age++
    fmt.Printf("%sさんの誕生日です！%d歳になりました。\n", p.Name, p.Age)
}

// インターフェースの定義
type Speaker interface {
    Speak() string
}

// 構造体にインターフェースを実装
func (p Person) Speak() string {
    return fmt.Sprintf("%sが話しています", p.Name)
}

func main() {
    // 構造体の初期化
    person1 := Person{
        Name: "田中太郎",
        Age:  25,
        City: "東京",
    }
    
    // メソッドの呼び出し
    person1.Introduce()
    person1.HaveBirthday()
    
    // インターフェースの使用
    var speaker Speaker = person1
    fmt.Println(speaker.Speak())
    
    // 構造体のポインタ
    person2 := &Person{"佐藤花子", 30, "大阪"}
    person2.Introduce()
}
