package models

import (
	"log"
	"time"
)

// Userのstructを定義
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

// Userを作成する関数
func (u *User) CreateUser() (err error) {
	// userを作成するコマンドを定義
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	// insert文を実行（このファイル内ではDbは定義していないが、modelsの中にあるから使用できる）
	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
