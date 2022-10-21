package logger

type Logger interface {
	LogError(err error, format string, args ...interface{})
	LogFatal(err error, format string, args ...interface{})
	LogInfo(format string, args ...interface{})
	LogWarn(format string, args ...interface{})
	Printf(format string, values ...interface{})
}
