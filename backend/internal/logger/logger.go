package logger

//imports for zap logging library and standard libraries
import (
	"fmt"
	"os"
	"strings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

)

// Logger is a wrapper around zap logger 
type Logger struct {
	*zap.Logger
	sugar *zap.SugaredLogger
}

var globalLogger *Logger

// initialze the global LOgger
func InitLogger(loglevel string , format string) (*Logger , error) {
	level := zapcore.InfoLevel
	switch strings.ToLower(logLevel) {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevl
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel		
	}
    var config zap.Config 
	if strings.ToLower(format) == "json" {
		config = zap.NewProductionConfig()
        config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO08601TimeEncoder

	} 
	else {
		config = zap.NewDevelopmentConfig()

	}

	config.level = zap.NewAtomicLevelAt(level)
	config.OutputPaths = []string("stdout")
	config.ErrorOutputPaths = []string("stderr")

	zapLoger,err :=config.Build()
	if err != nil {
		return nil , fmt.Errorf("failed to create logger %w",err)

	}
	logger :=&Logger{
		Logger : zapLogger,
		sugar : zapLogger.Sugar(),

	}

	globalLogger = logger
	return logger,nil
}

//functions for logging 
 func GetLogs() *Logger {
	if globalLogger == nil {
		var err error
		gloabalLogger , err = InitLogger("info","json")
		if err !=nil {
			fmt.Fprintf(os.Stderr,"Failed to initialize logger:%v\n",err)
			os.Exit(1)
		}
	}
	return globalLogger

 }

// Denug logs a debug message
func (l *Logger) Debug(msg string,fields ...zap.Field) {
	l.Logger.Debug(msg,fields...)

}
func (l *Logger) Infor(msg string,fields ...zap.Field) {
	l.Logger.Info(msg,fields...)

}

// warn logs a warnng message
func (l *Logger) Warn(msg string,fields ...zap.Field) {
	l.Logger.Warm(msg,fields...)

}

// Error logs an error message
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.Logger.Error(msg, fields...)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.Logger.Fatal(msg, fields...)
}

// Debugf logs a debug message using format string
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sugar.Debugf(template, args...)
}

// Infof logs an info message using format string
func (l *Logger) Infof(template string, args ...interface{}) {
	l.sugar.Infof(template, args...)
}

// Warnf logs a warning message using format string
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.sugar.Warnf(template, args...)
}

// Errorf logs an error message using format string
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.sugar.Errorf(template, args...)
}

// Fatalf logs a fatal message using format string and exits
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.sugar.Fatalf(template, args...)
}

// WithField returns a new logger with an additional field
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{
		Logger: l.Logger.With(zap.Any(key, value)),
		sugar:  l.sugar.With(key, value),
	}
}

// WithFields returns a new logger with additional fields
func (l *Logger) WithFields(fields ...zap.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
		sugar:  l.sugar,
	}
}

// Sync flushes any buffered log entries
func (l *Logger) Sync() error {
	return l.Logger.Sync()
}

