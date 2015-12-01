package logger

import (
	"log"
	"os"

	. "gopkg.in/check.v1"
)

type BuiltinTestSuite struct {
}

var (
	_ = Suite(&LevelTestSuite{})
)

func ExampleBuiltinDebug() {
	logger := NewBuiltinLogger(log.New(os.Stdout, "", 0), LevelDebug)
	logger.Debug("message1", "message2")

	// Output:
	// [DEBUG] message1message2
}

func ExampleBuiltinDebugf() {
	logger := NewBuiltinLogger(log.New(os.Stdout, "", 0), LevelDebug)
	logger.Debugf("debug messages: %s, %s", "message1", "message2")

	// Output:
	// [DEBUG] debug messages: message1, message2
}

func ExampleBuiltinDebugDisabled() {
	logger := NewBuiltinLogger(log.New(os.Stdout, "", 0), LevelError)
	logger.Debug("message")

	// Output:
}

func ExampleBuiltinStderr() {
	logger := NewBuiltinLogger(log.New(os.Stderr, "prefix", 0), LevelDebug)
	logger.Alert("message")

	// Error:
	// prefix [ALERT] message
}

func (s *BuiltinTestSuite) TestIsLevelEnabled(c *C) {
	logger := NewBuiltinLogger(log.New(os.Stdout, "", 0), LevelError)
	c.Assert(logger.IsDebugEnabled(), Equals, false)
	c.Assert(logger.IsInfoEnabled(), Equals, false)
	c.Assert(logger.IsNoticeEnabled(), Equals, false)
	c.Assert(logger.IsWarningEnabled(), Equals, false)
	c.Assert(logger.IsErrorEnabled(), Equals, true)
	c.Assert(logger.IsAlertEnabled(), Equals, true)
	c.Assert(logger.IsCriticalEnabled(), Equals, true)
	c.Assert(logger.IsEmergencyEnabled(), Equals, true)

}
