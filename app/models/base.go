package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

// テーブル名を定義
const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	// テーブルがない場合は作成する定義(users)
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	// テーブルのCREATE文を実行(users)
	Db.Exec(cmdU)

	// テーブルがない場合は作成するコマンドを定義(todos)
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	// テーブルのCREATE文を実行(todos)
	Db.Exec(cmdT)
}

// userのUUIDを作成する関数
func createUUID() (uuidobj uuid.UUID) {
	// UUIDを作成
	uuidobj, _ = uuid.NewUUID()

	// UUIDを返す
	return uuidobj
}

// passwordをハッシュ化する関数
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
