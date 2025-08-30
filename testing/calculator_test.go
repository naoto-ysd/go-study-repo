package calc

import (
	"testing"
)

// テスト関数は Test で始まる必要がある
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2

	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(4, 5)
	expected := 20

	if result != expected {
		t.Errorf("Multiply(4, 5) = %d; expected %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	// 正常ケース
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) でエラーが発生: %v", err)
	}

	expected := 5
	if result != expected {
		t.Errorf("Divide(10, 2) = %d; expected %d", result, expected)
	}

	// エラーケース（ゼロ除算）
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) でエラーが発生しませんでした")
	}
}

// ベンチマークテスト
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(123, 456)
	}
}

func TestMod(t *testing.T) {
	// 正常ケース
	result, err := Mod(10, 3)
	if err != nil {
		t.Fatalf("Mod(10, 3) でエラーが発生: %v", err)
	}
	if result != 1 {
		t.Errorf("Mod(10, 3) = %d; expected %d", result, 1)
	}

	// エラーケース（ゼロ除算）
	_, err = Mod(10, 0)
	if err == nil {
		t.Error("Mod(10, 0) でエラーが発生しませんでした")
	}
}
