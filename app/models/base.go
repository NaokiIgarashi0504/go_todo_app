package models

import (
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

// テーブル名を定義
const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	// usersテーブルがない場合は定義
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	// usersテーブルのCREATE文を実行
	Db.Exec(cmdU)
}
