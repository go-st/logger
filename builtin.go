package logger

import (
	"fmt"
	"log"
)

// BuiltinLogger is an adapter for built-in logger
type BuiltinLogger struct {
	logger *log.Logger
	level  Level
}

// NewBuiltinLogger creates new BuiltinLogger
func NewBuiltinLogger(logger *log.Logger, level Level) *BuiltinLogger {
	return &BuiltinLogger{logger: logger, level: level}
}

func (l *BuiltinLogger) Debugf(format string, args ...interface{}) {
	l.Logf(LevelDebug, format, args...)
}

func (l *BuiltinLogger) Infof(format string, args ...interface{}) {
	l.Logf(LevelInfo, format, args...)
}

func (l *BuiltinLogger) Noticef(format string, args ...interface{}) {
	l.Logf(LevelNotice, format, args...)
}

func (l *BuiltinLogger) Warningf(format string, args ...interface{}) {
	l.Logf(LevelWarning, format, args...)
}

func (l *BuiltinLogger) Errorf(format string, args ...interface{}) {
	l.Logf(LevelError, format, args...)
}

func (l *BuiltinLogger) Alertf(format string, args ...interface{}) {
	l.Logf(LevelAlert, format, args...)
}

func (l *BuiltinLogger) Criticalf(format string, args ...interface{}) {
	l.Logf(LevelCritical, format, args...)
}

func (l *BuiltinLogger) Emergencyf(format string, args ...interface{}) {
	l.Logf(LevelEmergency, format, args...)
}

func (l *BuiltinLogger) Logf(level Level, format string, args ...interface{}) {
	l.logMessage(level, fmt.Sprintf(format, args...))
}

func (l *BuiltinLogger) Debug(args ...interface{}) {
	l.Log(LevelDebug, args...)
}

func (l *BuiltinLogger) Info(args ...interface{}) {
	l.Log(LevelInfo, args...)
}

func (l *BuiltinLogger) Notice(args ...interface{}) {
	l.Log(LevelNotice, args...)
}

func (l *BuiltinLogger) Warning(args ...interface{}) {
	l.Log(LevelWarning, args...)
}

func (l *BuiltinLogger) Error(args ...interface{}) {
	l.Log(LevelError, args...)
}

func (l *BuiltinLogger) Alert(args ...interface{}) {
	l.Log(LevelAlert, args...)
}

func (l *BuiltinLogger) Critical(args ...interface{}) {
	l.Log(LevelCritical, args...)
}

func (l *BuiltinLogger) Emergency(args ...interface{}) {
	l.Log(LevelEmergency, args...)
}

func (l *BuiltinLogger) Log(level Level, args ...interface{}) {
	l.logMessage(level, fmt.Sprint(args...))
}

func (l *BuiltinLogger) IsDebugEnabled() bool {
	return l.isLevelEnabled(LevelDebug)
}

func (l *BuiltinLogger) IsInfoEnabled() bool {
	return l.isLevelEnabled(LevelInfo)
}

func (l *BuiltinLogger) IsNoticeEnabled() bool {
	return l.isLevelEnabled(LevelNotice)
}

func (l *BuiltinLogger) IsWarningEnabled() bool {
	return l.isLevelEnabled(LevelWarning)
}

func (l *BuiltinLogger) IsErrorEnabled() bool {
	return l.isLevelEnabled(LevelError)
}

func (l *BuiltinLogger) IsAlertEnabled() bool {
	return l.isLevelEnabled(LevelAlert)
}

func (l *BuiltinLogger) IsCriticalEnabled() bool {
	return l.isLevelEnabled(LevelCritical)
}

func (l *BuiltinLogger) IsEmergencyEnabled() bool {
	return l.isLevelEnabled(LevelEmergency)
}

func (l *BuiltinLogger) isLevelEnabled(level Level) bool {
	return level <= l.level
}

func (l *BuiltinLogger) logMessage(level Level, message string) {
	if l.isLevelEnabled(level) {
		l.logger.Print(fmt.Sprintf("[%s] %s", level, message))
	}
}
