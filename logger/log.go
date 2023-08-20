package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLog(logPath string) {
	encoder := getEncoder()
	sync := getWriteSync(logPath)
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	Logger = zap.New(core)
}

// set encoding format json
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

// set log path
func getWriteSync(logPath string) zapcore.WriteSyncer {
	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
