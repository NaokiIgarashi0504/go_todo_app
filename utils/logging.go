package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// 読み書き、ファイルの追加、追記する設定にして、変数に代入
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	// エラーの場合
	if err != nil {
		// ログを出力
		log.Fatalln(err)
	}

	// 書き込み先を標準出力とログファイルに指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)

	// フォーマット指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
