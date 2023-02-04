package controllers

import (
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
	_, err := session(w, r)

	if err != nil {
		// エラーの場合
		http.Redirect(w, r, "/", 302)
	} else {
		// セッションが正しい場合は、indexを表示
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
