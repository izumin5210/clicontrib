package clog

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger  *zap.Logger
	sLogger *zap.SugaredLogger
)

func init() {
	initialize(logModeNop)
}

// Set sets the logger used in application globa.
func Set(l *zap.Logger) {
	logger = l
	sLogger = logger.Sugar()
}

// InitDebug initialize the logger in debug mode.
func InitDebug() {
	initialize(logModeDebug)
}

// InitVerbose initialize the logger in verbose mode.
func InitVerbose() {
	initialize(logModeVerbose)
}

func initialize(m logMode) {
	Set(create(m).WithOptions(zap.AddCallerSkip(1)))
}

func create(m logMode) (l *zap.Logger) {
	var cfg zap.Config
	opts := []zap.Option{}
	l = zap.NewNop()

	switch m {
	case logModeVerbose:
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Local().Format("2006-01-02 15:04:05 MST"))
		}
		opts = append(opts, zap.AddStacktrace(zapcore.ErrorLevel))
	case logModeDebug:
		cfg = zap.NewProductionConfig()
		cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		cfg.DisableStacktrace = true
	default:
		return
	}

	logger, err := cfg.Build(opts...)
	if err == nil {
		l = logger
	}
	return
}

// Logger returns current logger object
func Logger() *zap.Logger {
	return logger
}

// Close closes the logger
func Close() error {
	var errs []error
	if err := logger.Sync(); err != nil {
		errs = append(errs, err)
	}
	if err := sLogger.Sync(); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

// Debug logs message and key-values pairs as DEBUG
func Debug(msg string, args ...interface{}) {
	sLogger.Debugw(msg, args...)
}

// Info logs message and key-values pairs as INFO
func Info(msg string, args ...interface{}) {
	sLogger.Infow(msg, args...)
}

// Warn logs message and key-values pairs in as WARN
func Warn(msg string, args ...interface{}) {
	sLogger.Warnw(msg, args...)
}

// Error logs message and key-values pairs in as ERROR
func Error(msg string, args ...interface{}) {
	sLogger.Errorw(msg, args...)
}
