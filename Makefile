.PHONY: help run-basics run-structs run-concurrency run-errors run-todo test clean

# デフォルトターゲット
help:
	@echo "Go学習リポジトリ - 使用可能なコマンド:"
	@echo "  make run-basics     - 基礎編のサンプルを実行"
	@echo "  make run-structs    - 構造体とメソッドのサンプルを実行"
	@echo "  make run-concurrency - 並行処理のサンプルを実行"
	@echo "  make run-errors     - エラーハンドリングのサンプルを実行"
	@echo "  make run-todo       - TodoCLIアプリケーションを実行"
	@echo "  make test           - テストを実行"
	@echo "  make clean          - ビルドファイルを削除"

# 基礎編の実行
run-basics:
	@echo "=== Hello World ==="
	go run basics/01_hello_world.go
	@echo "\n=== 制御構文 ==="
	go run basics/03_control_flow.go

# 変数と関数の実行
run-variables:
	@echo "\n=== 変数と定数 ==="
	go run basics/02_variables.go
	
# 構造体とメソッドの実行
run-structs:
	@echo "=== 構造体とメソッド ==="
	go run structs/01_structs_and_methods.go

# 並行処理の実行
run-concurrency:
	@echo "=== ゴルーチンとチャネル ==="
	go run concurrency/01_goroutines.go

# エラーハンドリングの実行
run-errors:
	@echo "=== エラーハンドリング ==="
	go run error-handling/01_error_handling.go

# TodoCLIアプリケーションの実行
run-todo:
	@echo "=== Todo CLIアプリケーション ==="
	go run projects/todo_cli.go

# テストの実行
test:
	@echo "=== テスト実行 ==="
	go test ./testing -v
	@echo "\n=== ベンチマークテスト ==="
	go test ./testing -bench=.

# クリーンアップ
clean:
	go clean
	rm -f *.exe
