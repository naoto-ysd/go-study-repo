package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// Todoアイテムの構造体
type Todo struct {
    ID        int
    Text      string
    Completed bool
}

// TodoListの管理
type TodoList struct {
    todos  []Todo
    nextID int
}

// 新しいTodoを追加
func (tl *TodoList) AddTodo(text string) {
    todo := Todo{
        ID:        tl.nextID,
        Text:      text,
        Completed: false,
    }
    tl.todos = append(tl.todos, todo)
    tl.nextID++
    fmt.Printf("Todo追加: %s\n", text)
}

// Todoリストを表示
func (tl *TodoList) ListTodos() {
    if len(tl.todos) == 0 {
        fmt.Println("Todoはありません")
        return
    }
    
    fmt.Println("=== Todoリスト ===")
    for _, todo := range tl.todos {
        status := "未完了"
        if todo.Completed {
            status = "完了"
        }
        fmt.Printf("%d: %s [%s]\n", todo.ID, todo.Text, status)
    }
}

// Todoを完了にマーク
func (tl *TodoList) CompleteTodo(id int) {
    for i := range tl.todos {
        if tl.todos[i].ID == id {
            tl.todos[i].Completed = true
            fmt.Printf("Todo %d を完了にマークしました\n", id)
            return
        }
    }
    fmt.Printf("ID %d のTodoが見つかりません\n", id)
}

func main() {
    todoList := &TodoList{nextID: 1}
    scanner := bufio.NewScanner(os.Stdin)
    
    fmt.Println("=== 簡単なTodo CLI ===")
    fmt.Println("コマンド: add <テキスト>, list, complete <ID>, quit")
    
    for {
        fmt.Print("> ")
        if !scanner.Scan() {
            break
        }
        
        input := strings.TrimSpace(scanner.Text())
        parts := strings.Fields(input)
        
        if len(parts) == 0 {
            continue
        }
        
        command := parts[0]
        
        switch command {
        case "add":
            if len(parts) < 2 {
                fmt.Println("使用方法: add <テキスト>")
                continue
            }
            text := strings.Join(parts[1:], " ")
            todoList.AddTodo(text)
            
        case "list":
            todoList.ListTodos()
            
        case "complete":
            if len(parts) < 2 {
                fmt.Println("使用方法: complete <ID>")
                continue
            }
            id, err := strconv.Atoi(parts[1])
            if err != nil {
                fmt.Println("無効なID:", parts[1])
                continue
            }
            todoList.CompleteTodo(id)
            
        case "quit":
            fmt.Println("終了します")
            return
            
        default:
            fmt.Println("不明なコマンド:", command)
            fmt.Println("使用可能なコマンド: add, list, complete, quit")
        }
    }
}
