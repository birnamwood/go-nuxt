package database

import (
	"testing"

	"github.com/birnamwood/go-nuxt/initialize/config"
)

//TestInit データベース初期化テスト。Docker上のDBは見られないのでエラーになる
func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "database初期化テスト",
		},
	}
	for _, tt := range tests {
		config.Init("development", "../config/environments/")
		t.Run(tt.name, func(t *testing.T) {
			Init()
		})
	}
}
