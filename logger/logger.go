package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	//make this variable private as we have create GetLogger function for making it exported
	log *zap.Logger
)

//in this init function we initialize the log
func init() {
	//configuration for how our log should come
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}

}

func GetLogger() *zap.Logger {
	return log
}

//we create a variadic function to not show our log library in every repo
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	log.Error(msg, tags...)
	log.Sync()
}
