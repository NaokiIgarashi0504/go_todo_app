package config

import (
	"log"
	"todo_app/utils"

	"gopkg.in/go-ini/ini.v1"
)

// ConfigListの構造体を作成
type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

// ConfigListを外部から呼び出せるようにグローバルに変数宣言
var Config ConfigList

// main関数よりも先に呼ばれる
func init() {
	// config.iniを読み込んで、ConfigListを設定する関数を呼ぶ
	LoadConfig()

	utils.LoggingSettings(Config.LogFile)
}

// config.iniを読み込んで、ConfigListを設定する関数
func LoadConfig() {
	// config.iniを読み込む
	cfg, err := ini.Load("config.ini")

	// エラーの場合
	if err != nil {
		// エラーをログに出力
		log.Fatalln(err)
	}

	// 読み込んだconfig.iniを変数に代入
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
