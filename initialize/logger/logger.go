package logger

import (
	"os"

	"github.com/birnamwood/go-nuxt/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() *zap.Logger {
	c := config.GetConfig()

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   c.GetString("log.filename"),
		MaxSize:    c.GetInt("log.maxsize"),
		MaxBackups: c.GetInt("log.maxbackups"),
		MaxAge:     c.GetInt("log.maxage"),
	})

	e := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(e),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		zap.InfoLevel,
	)

	logger := zap.New(core, zap.AddCaller(), zap.Development())
	return logger
}
