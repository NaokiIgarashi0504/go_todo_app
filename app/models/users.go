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

// Sessionのstructを定義
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
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

// Userを更新する関数
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

// Userを削除する関数
func (u *User) DeleteUser() (err error) {
	// deleteコマンドの定義
	cmd := `delete from users where id = ?`

	// deleteコマンドの実行
	_, err = Db.Exec(cmd, u.ID)

	// エラーハンドリング
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// 入力されたemailからユーザー情報を取得する関数
func GetUserByEmail(email string) (user User, err error) {
	// userを定義
	user = User{}

	// ユーザー情報を取得するコマンドを定義
	cmd := `select id, uuid, name, email, password, created_at from users where email = ?`

	// ユーザー情報を取得するコマンドを実行
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)

	return
}

// セッションを作成する関数
func (u *User) CreateSession() (session Session, err error) {
	// セッションと定義
	session = Session{}

	// セッションを登録するコマンドを定義
	cmd1 := `insert into sessions(
		uuid,
		email,
		user_id,
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())

	// エラーハンドリング
	if err != nil {
		log.Panicln(err)
	}

	// 取得するコマンドを定義
	cmd2 := `select id, uuid, email, user_id, created_at from sessions where user_id = ? and email = ?`

	// 取得するコマンドの実行
	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err
}

// セッションに存在するかチェックする関数
func (sess *Session) CheckSession() (valid bool, err error) {
	// セッションの情報を取得するコマンドを定義
	cmd := `select id, uuid, email, user_id, created_at from sessions where uuid = ?`

	// セッションの情報を取得するコマンドの実行
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	// エラーの場合はセッションが存在しない
	if err != nil {
		valid = false
		return
	}

	// セッションのUUIDが初期値でない場合は、セッションが存在する
	if sess.ID != 0 {
		valid = true
	}

	return valid, err
}
