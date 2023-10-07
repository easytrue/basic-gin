package conf

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {

	logMode := zapcore.DebugLevel
	if !viper.GetBool("app.dev") {
		logMode = zapcore.InfoLevel
	}
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)
	// logger:=zap.New(core).Sugar() logger.Info("Error Msg: %s", err.Error())
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer() zapcore.WriteSyncer {
	stringSeparator := string(filepath.Separator)
	stringRootDir, _ := os.Getwd()
	stringLogFilePath := stringRootDir + stringSeparator + "log" + stringSeparator + time.Now().Format(time.DateOnly) + ".txt"
	// 日志写入指针
	lumberjackSync := &lumberjack.Logger{
		Filename:   stringLogFilePath,
		MaxSize:    viper.GetInt("log.maxSize"),
		MaxBackups: viper.GetInt("log.maxBackups"),
		MaxAge:     viper.GetInt("log.maxAge"),
		Compress:   viper.GetBool("log.compress"),
	}
	return zapcore.AddSync(lumberjackSync)
}
