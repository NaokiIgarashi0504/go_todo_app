package controllers

import (
	"net/http"
	"todo_app/config"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URLの登録（「/」に訪れたらtopを表示する）
	http.HandleFunc("/", top)

	// Webサーバー構築
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
