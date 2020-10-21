package config

import (
	"github.com/spf13/viper"
)

var c *viper.Viper

// Init initializes config
func Init(env string) {
	//Viperという設定ファイル用ライブラリを使う
	c = viper.New()
	//上から設定ファイル拡張子・ファイル名・パスを指定
	c.SetConfigFile("yaml")
	c.SetConfigName(env)
	c.AddConfigPath("config/environments/")
	if err := c.ReadInConfig(); err != nil {
		panic(err)
	}
}

// GetConfig returns config
func GetConfig() *viper.Viper {
	return c
}
