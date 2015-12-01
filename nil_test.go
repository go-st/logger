package logger

func ExampleNilDebug() {
	logger := NewNilLogger()
	logger.Debug("message1", "message2")

	// Output:
}

func ExampleNilDebugf() {
	logger := NewNilLogger()
	logger.Debugf("debug message: %s", "test")

	// Output:
}
