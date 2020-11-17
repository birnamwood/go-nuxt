package function

import (
	"io"

	"go.uber.org/zap"
)

//CloseFile 設定ファイルのClose時にエラーハンドリングを行う
func CloseFile(c io.Closer, f string) {
	if err := c.Close(); err != nil {
		zap.S().Error(f + "のCloseに失敗しました。")
	}
}
