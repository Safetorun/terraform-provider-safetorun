package logger

type MockLogger struct {
}

func (l MockLogger) Errorf(message string, args ...interface{}) {
}

func (l MockLogger) Debugf(message string, args ...interface{}) {
}

func (l MockLogger) Infof(message string, args ...interface{}) {
}

func (l MockLogger) Warnf(message string, args ...interface{}) {
}
