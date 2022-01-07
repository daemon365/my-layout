package logging

import (
	"os"

	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(filepath string) log.Logger {
	writeSyncer := getLogWriter(filepath)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	return kratoszap.NewLogger(logger)
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter(filepath string) zapcore.WriteSyncer {
	file, _ := os.Create(filepath)
	return zapcore.AddSync(file)
}
