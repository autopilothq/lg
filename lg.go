package lg

import (
	"os"
)

func Extend(f ...F) Log {
	fields := &Fields{}
	ext := &ExtendedLog{fields: fields}
	for _, fld := range f {
		fields.set(fld)
	}
	return ext
}

// Trace logs a message at trace level
func Trace(args ...interface{}) {
	addEntry(LevelTrace, args)
}

// Traceln logs a message at trace level
func Traceln(args ...interface{}) {
	addEntry(LevelTrace, args)
}

// Tracef logs a formatted message at trace level
func Tracef(pattern string, args ...interface{}) {
	addFormattedEntry(LevelTrace, pattern, args)
}

// Debug logs a message at debug level
func Debug(args ...interface{}) {
	addEntry(LevelDebug, args)
}

// Debugln logs a message at debug level
func Debugln(args ...interface{}) {
	addEntry(LevelDebug, args)
}

// Debugf logs a formatted message at debug level
func Debugf(pattern string, args ...interface{}) {
	addFormattedEntry(LevelDebug, pattern, args)
}

// Info logs a message at info level
func Info(args ...interface{}) {
	addEntry(LevelInfo, args)
}

// Infoln logs a message at info level
func Infoln(args ...interface{}) {
	addEntry(LevelInfo, args)
}

// Infof logs a formatted message at info level
func Infof(pattern string, args ...interface{}) {
	addFormattedEntry(LevelInfo, pattern, args)
}

// Print logs a message at info level
func Print(args ...interface{}) {
	addEntry(LevelInfo, args)
}

// Println logs a message at info level
func Println(args ...interface{}) {
	addEntry(LevelInfo, args)
}

// Printf logs a formatted message at info level
func Printf(pattern string, args ...interface{}) {
	addFormattedEntry(LevelInfo, pattern, args)
}

// Warn logs a message at warn level
func Warn(args ...interface{}) {
	addEntry(LevelWarn, args)
}

// Warnln logs a message at warn level
func Warnln(args ...interface{}) {
	addEntry(LevelWarn, args)
}

// Warnf logs a formatted message at warn level
func Warnf(pattern string, args ...interface{}) {
	addFormattedEntry(LevelWarn, pattern, args)
}

// Warning logs a message at warn level
func Warning(args ...interface{}) {
	addEntry(LevelWarn, args)
}

// Warningln logs a message at warn level
func Warningln(args ...interface{}) {
	addEntry(LevelWarn, args)
}

// Warningf logs a formatted message at warn level
func Warningf(pattern string, args ...interface{}) {
	addFormattedEntry(LevelWarn, pattern, args)
}

// Error logs a message at error level
func Error(args ...interface{}) {
	addEntry(LevelError, args)
}

// Errorln logs a message at error level
func Errorln(args ...interface{}) {
	addEntry(LevelError, args)
}

// Errorf logs a formatted message at error level
func Errorf(pattern string, args ...interface{}) {
	addFormattedEntry(LevelError, pattern, args)
}

// Fatal logs a message at fatal level
func Fatal(args ...interface{}) {
	addEntry(LevelFatal, args)
	os.Exit(1)
}

// Fatalln logs a message at fatal level
func Fatalln(args ...interface{}) {
	addEntry(LevelFatal, args)
	os.Exit(1)
}

// Fatalf logs a formatted message at fatal level
func Fatalf(pattern string, args ...interface{}) {
	addFormattedEntry(LevelFatal, pattern, args)
	os.Exit(1)
}

// Panic logs a message at fatal level and panics
func Panic(args ...interface{}) {
	entry, _ := addEntry(LevelFatal, args)
	panic(entry.Message)
}

// Panicln logs a message at fatal level and panics
func Panicln(args ...interface{}) {
	entry, _ := addEntry(LevelFatal, args)
	panic(entry.Message)
}

// Panicf logs a formatted message at fatal level and panics
func Panicf(pattern string, args ...interface{}) {
	entry, _ := addFormattedEntry(LevelFatal, pattern, args)
	panic(entry.Message)
}
