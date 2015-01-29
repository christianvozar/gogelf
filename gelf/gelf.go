// Copyright © 2014-2015, Civis Analytics

package gelf

// Level is a generic type for severity.
type Level int

// Levels matching RFC5424 severity.
const (
	LevelPanic Level = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInfo
	LevelDebug
)

// Panic generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 0.
// System is unusable.
func Panic(s, l string) {
	NewMessage(LevelPanic, s, l).Send()
}

// Alert generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 1.
// Action must be taken immediately.
func Alert(s, l string) {
	NewMessage(LevelAlert, s, l).Send()
}

// Crit generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 2.
// Critical conditions.
func Crit(s, l string) {
	NewMessage(LevelCritical, s, l).Send()
}

// Error generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 3.
// Error conditions.
func Error(s, l string) {
	NewMessage(LevelError, s, l).Send()
}

// Warn generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 4.
// Warning conditions.
func Warn(s, l string) {
	NewMessage(LevelWarning, s, l).Send()
}

// Notice generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 5.
// Normal but significant condition.
func Notice(s, l string) {
	NewMessage(LevelNotice, s, l).Send()
}

// Info generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 6.
// Informational messages.
func Info(s, l string) {
	NewMessage(LevelInfo, s, l).Send()
}

// Debug generates a Graylog2 Extended Format message with a RFC5424 severity
// code of 7.
// Developer debug messages.
func Debug(s, l string) {
	NewMessage(LevelDebug, s, l).Send()
}
