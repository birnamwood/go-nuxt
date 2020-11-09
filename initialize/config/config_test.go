package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestConfig(t *testing.T) {
	c = viper.New()
	c.SetConfigFile("yaml")
	c.SetConfigName("production")
	path, _ := os.Getwd()
	path2 := path + "/environments/"
	c.AddConfigPath(path2)

	if err := c.ReadInConfig(); err != nil {
		panic(err)
	}
	c := GetConfig()
	if c == nil {
		t.Error("Configの初期化に失敗しました。")
	}
}
