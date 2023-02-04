package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

// signup.htmlを表示する関数
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
