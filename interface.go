package logger

// ILogger is an common interface for loggers, which pretend to support RFC 5424 (https://tools.ietf.org/html/rfc5424).
type ILogger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Noticef(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Alertf(format string, args ...interface{})
	Criticalf(format string, args ...interface{})
	Emergencyf(format string, args ...interface{})
	Logf(level Level, format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Notice(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Alert(args ...interface{})
	Critical(args ...interface{})
	Emergency(args ...interface{})
	Log(level Level, args ...interface{})
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsNoticeEnabled() bool
	IsWarningEnabled() bool
	IsErrorEnabled() bool
	IsAlertEnabled() bool
	IsCriticalEnabled() bool
	IsEmergencyEnabled() bool
}
