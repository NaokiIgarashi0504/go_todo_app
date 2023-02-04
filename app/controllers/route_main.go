package controllers

import (
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	// セッションの確認
	_, err := session(w, r)

	if err != nil {
		// ログインしていない場合は、public_navbarを表示
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		// ログインしている場合は、todosを表示
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// セッションの確認
	sess, err := session(w, r)

	if err != nil {
		// エラーの場合
		http.Redirect(w, r, "/", 302)
	} else {
		// セッションが正しい場合
		// セッションからユーザーの情報を取得
		user, err := sess.GetUserBySession()

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		// ユーザーに紐づくtodoを取得
		todos, _ := user.GetTodosByUser()

		// ユーザーのtodosに代入
		user.Todos = todos

		// indexを表示（userの情報を渡す）
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}
