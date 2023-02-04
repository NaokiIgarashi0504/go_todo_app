package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	// templateの構造体を生成
	t, err := template.ParseFiles("app/views/templates/top.html")

	// エラーハンドリング
	if err != nil {
		log.Panicln(err)
	}
	t.Execute(w, "hello")
}
