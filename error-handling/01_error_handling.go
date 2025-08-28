package main

import (
    "errors"
    "fmt"
    "strconv"
)

// カスタムエラー型
type ValidationError struct {
    Field string
    Value string
    Msg   string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("検証エラー - フィールド: %s, 値: %s, メッセージ: %s", 
        e.Field, e.Value, e.Msg)
}

// 年齢を検証する関数
func validateAge(ageStr string) (int, error) {
    age, err := strconv.Atoi(ageStr)
    if err != nil {
        return 0, fmt.Errorf("年齢の変換に失敗: %w", err)
    }
    
    if age < 0 {
        return 0, &ValidationError{
            Field: "age",
            Value: ageStr,
            Msg:   "年齢は0以上である必要があります",
        }
    }
    
    if age > 150 {
        return 0, &ValidationError{
            Field: "age",
            Value: ageStr,
            Msg:   "年齢は150以下である必要があります",
        }
    }
    
    return age, nil
}

// ファイルを読み込む関数（エラーの例）
func readFile(filename string) (string, error) {
    if filename == "" {
        return "", errors.New("ファイル名が空です")
    }
    
    // 実際のファイル読み込みはしないが、エラーの例として
    if filename == "notfound.txt" {
        return "", fmt.Errorf("ファイル '%s' が見つかりません", filename)
    }
    
    return "ファイルの内容", nil
}

func main() {
    fmt.Println("=== エラーハンドリングの例 ===")
    
    // 正常なケース
    age, err := validateAge("25")
    if err != nil {
        fmt.Printf("エラー: %v\n", err)
    } else {
        fmt.Printf("年齢: %d\n", age)
    }
    
    // 数値変換エラー
    _, err = validateAge("abc")
    if err != nil {
        fmt.Printf("エラー: %v\n", err)
    }
    
    // カスタムエラー（負の数）
    _, err = validateAge("-5")
    if err != nil {
        fmt.Printf("エラー: %v\n", err)
        
        // エラーの型をチェック
        var validationErr *ValidationError
        if errors.As(err, &validationErr) {
            fmt.Printf("検証エラーの詳細 - フィールド: %s\n", validationErr.Field)
        }
    }
    
    // ファイル読み込みエラー
    content, err := readFile("notfound.txt")
    if err != nil {
        fmt.Printf("ファイル読み込みエラー: %v\n", err)
    } else {
        fmt.Printf("ファイル内容: %s\n", content)
    }
    
    // 正常なファイル読み込み
    content, err = readFile("example.txt")
    if err != nil {
        fmt.Printf("エラー: %v\n", err)
    } else {
        fmt.Printf("ファイル内容: %s\n", content)
    }
}
