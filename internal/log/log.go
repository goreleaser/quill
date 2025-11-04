/*
Package log contains the singleton object and helper functions for facilitating logging within the library.
*/
package log

func Set(l any) {}
func Get() any  { return nil }

// Errorf takes a formatted template string and template arguments for the error logging level.
func Errorf(format string, args ...interface{}) {
}

// Error logs the given arguments at the error logging level.
func Error(args ...interface{}) {
}

// Warnf takes a formatted template string and template arguments for the warning logging level.
func Warnf(format string, args ...interface{}) {
}

// Warn logs the given arguments at the warning logging level.
func Warn(args ...interface{}) {
}

// Infof takes a formatted template string and template arguments for the info logging level.
func Infof(format string, args ...interface{}) {
}

// Info logs the given arguments at the info logging level.
func Info(args ...interface{}) {
}

// Debugf takes a formatted template string and template arguments for the debug logging level.
func Debugf(format string, args ...interface{}) {
}

// Debug logs the given arguments at the debug logging level.
func Debug(args ...interface{}) {
}

// Tracef takes a formatted template string and template arguments for the trace logging level.
func Tracef(format string, args ...interface{}) {
}

// Trace logs the given arguments at the trace logging level.
func Trace(args ...interface{}) {
}

// WithFields returns a message logger with multiple key-value fields.
func WithFields(fields ...interface{}) Dummy {
	return Dummy{}
}

type Dummy struct{}

// Errorf takes a formatted template string and template arguments for the error logging level.
func (Dummy) Errorf(format string, args ...interface{}) {
}

// Error logs the given arguments at the error logging level.
func (Dummy) Error(args ...interface{}) {
}

// Warnf takes a formatted template string and template arguments for the warning logging level.
func (Dummy) Warnf(format string, args ...interface{}) {
}

// Warn logs the given arguments at the warning logging level.
func (Dummy) Warn(args ...interface{}) {
}

// Infof takes a formatted template string and template arguments for the info logging level.
func (Dummy) Infof(format string, args ...interface{}) {
}

// Info logs the given arguments at the info logging level.
func (Dummy) Info(args ...interface{}) {
}

// Debugf takes a formatted template string and template arguments for the debug logging level.
func (Dummy) Debugf(format string, args ...interface{}) {
}

// Debug logs the given arguments at the debug logging level.
func (Dummy) Debug(args ...interface{}) {
}

// Tracef takes a formatted template string and template arguments for the trace logging level.
func (Dummy) Tracef(format string, args ...interface{}) {
}

// Trace logs the given arguments at the trace logging level.
func (Dummy) Trace(args ...interface{}) {
}

// WithFields returns a message logger with multiple key-value fields.
func (Dummy) WithFields(fields ...interface{}) Dummy {
	return Dummy{}
}
