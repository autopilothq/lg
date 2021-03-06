package lg

import (
	"os"
)

// Extend returns a new sub logger by extending the current one with
// extra fields.
func (e ExtendedLog) Extend(f ...F) Log {
	newFields := Fields{}

	for _, fld := range e.fields.contents {
		newFields.set(fld)
	}

	for _, fld := range f {
		newFields.set(fld)
	}

	return ExtendedLog{
		fields: newFields,
		prefix: e.prefix,
	}
}

// ExtendPrefix returns a new sub logger by extending the current one with
// prefix and extra fields.
func (e ExtendedLog) ExtendPrefix(prefix string, f ...F) Log {
	ext := e.Extend(f...).(ExtendedLog)
	if ext.prefix == "" {
		ext.prefix = prefix
	} else {
		ext.prefix += prefixDelimiter + prefix
	}
	return ext
}

// ExtendedLog is a concrete logger interface, optionally with specific fields
type ExtendedLog struct {
	prefix string
	fields Fields
}

func (e ExtendedLog) addEntry(level Level, args []interface{}) (*Entry, error) {
	if e.fields.contents != nil {
		i := 0
		newArgs := make([]interface{}, len(args)+len(e.fields.contents))
		for ; i < len(e.fields.contents); i++ {
			newArgs[i] = e.fields.contents[i]
		}
		for j := 0; j < len(args); i, j = i+1, j+1 {
			newArgs[i] = args[j]
		}
		return addEntry(level, e.prefix, newArgs)
	}
	return addEntry(level, e.prefix, args)
}

func (e ExtendedLog) addFormattedEntry(
	level Level, pattern string, args []interface{},
) (*Entry, error) {
	fields, remaining := ExtractTrailingFields(args)

	if e.fields.contents != nil {
		i := 0
		newArgs := make([]interface{},
			len(remaining)+len(e.fields.contents)+len(fields.contents))
		for ; i < len(remaining); i++ {
			newArgs[i] = remaining[i]
		}
		for j := 0; j < len(e.fields.contents); i, j = i+1, j+1 {
			newArgs[i] = e.fields.contents[j]
		}
		for k := 0; k < len(fields.contents); i, k = i+1, k+1 {
			newArgs[i] = fields.contents[k]
		}
		return addFormattedEntry(level, e.prefix, pattern, newArgs)
	}
	return addFormattedEntry(level, e.prefix, pattern, args)
}

// Trace logs a message at trace level
func (e ExtendedLog) Trace(args ...interface{}) {
	e.addEntry(LevelTrace, args)
}

// Traceln logs a message at trace level
func (e ExtendedLog) Traceln(args ...interface{}) {
	e.addEntry(LevelTrace, args)
}

// Tracef logs a formatted message at trace level
func (e ExtendedLog) Tracef(pattern string, args ...interface{}) {
	e.addFormattedEntry(LevelTrace, pattern, args)
}

// Debug logs a message at debug level
func (e ExtendedLog) Debug(args ...interface{}) {
	e.addEntry(LevelDebug, args)
}

// Debugln logs a message at debug level
func (e ExtendedLog) Debugln(args ...interface{}) {
	e.addEntry(LevelDebug, args)
}

// Debugf logs a formatted message at debug level
func (e ExtendedLog) Debugf(pattern string, args ...interface{}) {
	e.addFormattedEntry(LevelDebug, pattern, args)
}

// Info logs a message at info level
func (e ExtendedLog) Info(args ...interface{}) {
	e.addEntry(LevelInfo, args)
}

// Infoln logs a message at info level
func (e ExtendedLog) Infoln(args ...interface{}) {
	e.addEntry(LevelInfo, args)
}

// Infof logs a formatted message at info level
func (e ExtendedLog) Infof(pattern string, args ...interface{}) {
	e.addFormattedEntry(LevelInfo, pattern, args)
}

// Print logs a message at info level
func (e ExtendedLog) Print(args ...interface{}) {
	e.addEntry(LevelInfo, args)
}

// Println logs a message at info level
func (e ExtendedLog) Println(args ...interface{}) {
	e.addEntry(LevelInfo, args)
}

// Printf logs a formatted message at info level
func (e ExtendedLog) Printf(pattern string, args ...interface{}) {
	e.addFormattedEntry(LevelInfo, pattern, args)
}

// Warn logs a message at warn level
func (e ExtendedLog) Warn(args ...interface{}) {
	e.addEntry(LevelWarn, args)
}

// Warnln logs a message at warn level
func (e ExtendedLog) Warnln(args ...interface{}) {
	e.addEntry(LevelWarn, args)
}

// Warnf logs a formatted message at warn level
func (e ExtendedLog) Warnf(pattern string, args ...interface{}) {
	e.addFormattedEntry(LevelWarn, pattern, args)
}

// Warning logs a message at warn level
func (e ExtendedLog) Warning(args ...interface{}) {
	e.addEntry(LevelWarn, args)
}

// Warningln logs a message at warn level
func (e ExtendedLog) Warningln(args ...interface{}) {
	e.addEntry(LevelWarn, args)
}

// Warningf logs a formatted message at warn level
func (e ExtendedLog) Warningf(pattern string, args ...interface{}) {
	e.addFormattedEntry(LevelWarn, pattern, args)
}

// Error logs a message at error level
func (e ExtendedLog) Error(args ...interface{}) {
	e.addEntry(LevelError, args)
}

// Errorln logs a message at error level
func (e ExtendedLog) Errorln(args ...interface{}) {
	e.addEntry(LevelError, args)
}

// Errorf logs a formatted message at error level
func (e ExtendedLog) Errorf(pattern string, args ...interface{}) {
	e.addFormattedEntry(LevelError, pattern, args)
}

// Fatal logs a message at fatal level
func (e ExtendedLog) Fatal(args ...interface{}) {
	e.addEntry(LevelFatal, args)
	os.Exit(1)
}

// Fatalln logs a message at fatal level
func (e ExtendedLog) Fatalln(args ...interface{}) {
	e.addEntry(LevelFatal, args)
	os.Exit(1)
}

// Fatalf logs a formatted message at fatal level
func (e ExtendedLog) Fatalf(pattern string, args ...interface{}) {
	e.addFormattedEntry(LevelFatal, pattern, args)
	os.Exit(1)
}

// Panic logs a message at fatal level and panics
func (e ExtendedLog) Panic(args ...interface{}) {
	entry, _ := e.addEntry(LevelFatal, args)
	panic(entry.Message)
}

// Panicln logs a message at fatal level and panics
func (e ExtendedLog) Panicln(args ...interface{}) {
	entry, _ := e.addEntry(LevelFatal, args)
	panic(entry.Message)
}

// Panicf logs a formatted message at fatal level and panics
func (e ExtendedLog) Panicf(pattern string, args ...interface{}) {
	entry, _ := e.addFormattedEntry(LevelFatal, pattern, args)
	panic(entry.Message)
}
