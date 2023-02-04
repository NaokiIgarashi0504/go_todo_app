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

// データ取得（todo）
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

// データ取得（todos）
func GetTodos() (todos []Todo, err error) {
	// selectコマンドを定義
	cmd := `select id, content, user_id, created_at from todos`

	// selectコマンドの実行
	rows, err := Db.Query(cmd)

	// エラーハンドリング
	if err != nil {
		log.Panicln(err)
	}

	for rows.Next() {
		// 変数todoを定義
		var todo Todo

		// スキャン
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		// appendする
		todos = append(todos, todo)
	}

	rows.Close()

	return todos, err
}

// 特定のuserのtodoを取得する関数
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	// ユーザーに紐づくデータ取得のコマンドを定義
	cmd := `select id, content, user_id, created_at from todos where user_id = ?`

	// ユーザーに紐づくデータ取得のコマンドを実行
	rows, err := Db.Query(cmd, u.ID)

	// エラーハンドリング
	if err != nil {
		log.Panicln(err)
	}

	for rows.Next() {
		// 変数todoを定義
		var todo Todo

		// スキャン
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		// appendする
		todos = append(todos, todo)
	}

	rows.Close()

	return todos, err
}
