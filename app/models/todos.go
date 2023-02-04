package models

import (
	"log"
	"time"
)

// structを定義（todo）
type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

// 　todoを作成する関数
func (u *User) CreateTodo(content string) (err error) {
	// insertのコマンドを定義
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values (?, ?, ?)`

	// insertコマンドの実行
	_, err = Db.Exec(cmd, content, u.ID, time.Now())

	// エラーハンドリング
	if err != nil {
		log.Panicln(err)
	}

	return err
}

// データ取得（todos）
func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where id = ?`

	// 型の宣言
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}
