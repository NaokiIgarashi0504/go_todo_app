package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo_app/app/models"
	"todo_app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	// filesを定義
	var files []string

	// filenamesを回す
	for _, file := range filenames {
		// filesに代入
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// Mustはあらかじめtemplateをキャッシュしておいて効率良くする。ParseFilesは失敗の際にパニック状態になる。
	template := template.Must(template.ParseFiles(files...))

	// defineを使用している場合は、ExecuteTemplateを使用して明示的に宣言する必要がある
	template.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	// 設定しているcookieを取得
	cookie, err := r.Cookie("_cookie")

	// errが空の場合
	if err == nil {
		// 設定しているcookieの値を変数に代入
		sess = models.Session{UUID: cookie.Value}

		// セッションのチェック
		if ok, _ := sess.CheckSession(); !ok {
			// チェックで正しくなかった場合
			err = fmt.Errorf("Invalid session")
		}
	}

	return sess, err
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))

	// app/view内には動的ファイルが入る可能性があるため、静的ファイルであることを明示するためにstaticを使用
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URLの登録（「/」に訪れたらtopを表示する）
	http.HandleFunc("/", top)

	// URLの登録（「signup」に訪れたらtopを表示する）
	http.HandleFunc("/signup", signup)

	// URLの登録（「login」に訪れたらtopを表示する）
	http.HandleFunc("/login", login)

	// URLの登録（「authenticate」に訪れたらtopを表示する）
	http.HandleFunc("/authenticate", authenticate)

	// URLの登録（「todos」に訪れたらindexを表示する）（ログインしているユーザー）
	http.HandleFunc("/todos", index)

	// Webサーバー構築
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
