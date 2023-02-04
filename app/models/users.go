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

// Userを取得する関数
func GetUser(id int) (user User, err error) {
	// userを定義
	user = User{}

	// userの情報を取得するコマンドを定義
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`

	// userの情報を取得するコマンドの実行
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) UpdateUser() (err error) {
	// updateのコマンドを定義
	cmd := `update users set name = ?, email = ? where id = ?`

	// updateコマンドの実行
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)

	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
