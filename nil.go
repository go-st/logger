package logger

// NilLogger is a logger which does nothing. Feel free to use it in unit-tests.
type NilLogger struct{}

// NewNilLogger creates new NilLogger
func NewNilLogger() *NilLogger {
	return &NilLogger{}
}

func (l *NilLogger) Debugf(format string, args ...interface{}) {
}

func (l *NilLogger) Infof(format string, args ...interface{}) {
}

func (l *NilLogger) Noticef(format string, args ...interface{}) {
}

func (l *NilLogger) Warningf(format string, args ...interface{}) {
}

func (l *NilLogger) Errorf(format string, args ...interface{}) {
}

func (l *NilLogger) Alertf(format string, args ...interface{}) {
}

func (l *NilLogger) Criticalf(format string, args ...interface{}) {
}

func (l *NilLogger) Emergencyf(format string, args ...interface{}) {
}

func (l *NilLogger) Logf(level Level, format string, args ...interface{}) {
}

func (l *NilLogger) Debug(args ...interface{}) {
}

func (l *NilLogger) Info(args ...interface{}) {
}

func (l *NilLogger) Notice(args ...interface{}) {
}

func (l *NilLogger) Warning(args ...interface{}) {
}

func (l *NilLogger) Error(args ...interface{}) {
}

func (l *NilLogger) Alert(args ...interface{}) {
}

func (l *NilLogger) Critical(args ...interface{}) {
}

func (l *NilLogger) Emergency(args ...interface{}) {
}

func (l *NilLogger) Log(level Level, args ...interface{}) {
}

func (l *NilLogger) IsDebugEnabled() bool {
	return false
}

func (l *NilLogger) IsInfoEnabled() bool {
	return false
}

func (l *NilLogger) IsNoticeEnabled() bool {
	return false
}

func (l *NilLogger) IsWarningEnabled() bool {
	return false
}

func (l *NilLogger) IsErrorEnabled() bool {
	return false
}

func (l *NilLogger) IsAlertEnabled() bool {
	return false
}

func (l *NilLogger) IsCriticalEnabled() bool {
	return false
}

func (l *NilLogger) IsEmergencyEnabled() bool {
	return false
}
