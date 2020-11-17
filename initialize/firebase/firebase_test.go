package firebase

import "testing"

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		//引数返り値なし
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
		})
	}
}
