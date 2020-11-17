package config

import (
	"testing"
)

func TestInit(t *testing.T) {
	type args struct {
		env  string
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Development環境Configファイル生成",
			args: args{
				env:  "development",
				path: "./environments/",
			},
		},
		{
			name: "Production環境Configファイル生成",
			args: args{
				env:  "production",
				path: "./environments/",
			},
		},
		{
			name: "Test環境Configファイル生成",
			args: args{
				env:  "test",
				path: "./environments/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.env, tt.args.path)
		})
	}
}
