package controllers

import (
	"fmt"
	"html/template"
	"net/http"
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

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))

	// app/view内には動的ファイルが入る可能性があるため、静的ファイルであることを明示するためにstaticを使用
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URLの登録（「/」に訪れたらtopを表示する）
	http.HandleFunc("/", top)

	// Webサーバー構築
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
