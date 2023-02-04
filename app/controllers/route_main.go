package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
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

// 新たなtodoを作成する処理
func todoNew(w http.ResponseWriter, r *http.Request) {
	// セッションの確認
	_, err := session(w, r)

	if err != nil {
		// ログインしていない場合は、ログインページにリダイレクト
		http.Redirect(w, r, "/login", 302)
	} else {
		// ログインしている場合は、新たなtodoを作成するページを表示
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

// 新たなtodoを保存する処理
func todoSave(w http.ResponseWriter, r *http.Request) {
	// セッションの確認
	sess, err := session(w, r)

	if err != nil {
		// ログインしていない場合は、ログインページにリダイレクト
		http.Redirect(w, r, "/login", 302)
	} else {
		// ログインしている場合
		// 入力した値を取得
		err = r.ParseForm()

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		// ユーザー情報を取得
		user, err := sess.GetUserBySession()

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		// 入力した値を取得
		content := r.PostFormValue("content")

		// 新たなtodoを保存
		if err := user.CreateTodo(content); err != nil {
			log.Panicln(err)
		}

		// 新たなtodoの保存に成功したら、todo一覧ページにリダイレクト
		http.Redirect(w, r, "/todos", 302)
	}
}

// 編集に関する処理
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	// セッションの確認
	sess, err := session(w, r)

	if err != nil {
		// ログインしていない場合は、ログインページにリダイレクト
		http.Redirect(w, r, "/login", 302)
	} else {
		// ログインしている場合
		// ユーザー情報を取得
		_, err := sess.GetUserBySession()

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		// 該当のtodoを取得
		t, err := models.GetTodo(id)

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

// updateに関する処理
func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	// セッションの確認
	sess, err := session(w, r)

	if err != nil {
		// ログインしていない場合は、ログインページにリダイレクト
		http.Redirect(w, r, "/login", 302)
	} else {
		// ログインしている場合
		// フォームの内容を解析
		err := r.ParseForm()

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		// ユーザー情報を取得
		user, err := sess.GetUserBySession()

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		// contentの内容を取得
		content := r.PostFormValue("content")

		// ストラクトを作成
		t := &models.Todo{ID: id, Content: content, UserID: user.ID}

		// updateの実行
		if err := t.UpdateTodo(); err != nil {
			log.Panicln(err)
		}

		// 一覧画面にリダイレクト
		http.Redirect(w, r, "/todos", 302)
	}
}

// updateに関する処理
func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	// セッションの確認
	sess, err := session(w, r)

	if err != nil {
		// ログインしていない場合は、ログインページにリダイレクト
		http.Redirect(w, r, "/login", 302)
	} else {
		// ログインしている場合
		// ユーザー情報を取得
		_, err := sess.GetUserBySession()

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		// 該当のtodoを取得
		t, err := models.GetTodo(id)

		// エラーハンドリング
		if err != nil {
			log.Panicln(err)
		}

		// deleteの実行
		if err := t.DeleteTodo(); err != nil {
			log.Panicln(err)
		}

		// 一覧画面にリダイレクト
		http.Redirect(w, r, "/todos", 302)
	}
}
