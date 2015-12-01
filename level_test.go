package logger

import (
	. "gopkg.in/check.v1"
)

type LevelTestSuite struct {
}

var (
	_ = Suite(&LevelTestSuite{})
)

func (s *LevelTestSuite) TestString(c *C) {
	levels := map[Level]string{
		LevelEmergency: "EMERGENCY",
		LevelAlert:     "ALERT",
		LevelCritical:  "CRITICAL",
		LevelError:     "ERROR",
		LevelWarning:   "WARNING",
		LevelNotice:    "NOTICE",
		LevelInfo:      "INFO",
		LevelDebug:     "DEBUG",
		Level(10):      "UNKNOWN",
	}

	for level, name := range levels {
		c.Assert(level.String(), Equals, name)
	}
}
