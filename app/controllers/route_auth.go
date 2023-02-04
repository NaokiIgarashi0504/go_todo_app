package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

// サインアップに関する処理
func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// GETの場合
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		// POSTの場合
		// ParseFormでformの内容を解析できる
		err := r.ParseForm()

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}
	}

	// ユーザーのストラクトを作成
	user := models.User{
		// PostFormValueで引数を指定することで、該当のformの値を受け取れる
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		PassWord: r.PostFormValue("password"),
	}

	// ユーザーを作成
	if err := user.CreateUser(); err != nil {
		log.Panicln(err)
	}

	// ユーザーの登録に成功した場合はトップページにリダイレクト
	http.Redirect(w, r, "/", 302)
}

// ログイン画面の表示に関する処理
func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "login")
}

// ユーザーの認証に関する処理
func authenticate(w http.ResponseWriter, r *http.Request) {
	// フォームの情報を取得
	err := r.ParseForm()

	// エラーハンドリング
	if err != nil {
		log.Panicln(err)
	}

	// emailの情報をからユーザー情報を取得する
	user, err := models.GetUserByEmail(r.PostFormValue("email"))

	// エラーハンドリング
	if err != nil {
		log.Panicln(err)
		http.Redirect(w, r, "/login", 302)
	}

	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		// パスワードが正しい場合
		// セッションを作成
		session, err := user.CreateSession()

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		// cookieの作成
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}

		// cookieをセット
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else {
		// パスワードが一致しない場合
		http.Redirect(w, r, "/login", 302)
	}
}
