package lg

import (
	"bytes"
	"sync"
)

type MockLog struct {
	prefix  string
	fields  Fields
	entries []*Entry
	mutex   sync.RWMutex
	parent  *MockLog
}

type filter interface {
	check(*Entry) bool
}

func Mock() *MockLog {
	return &MockLog{}
}

func (m *MockLog) Count(filters ...filter) (count int) {
	if m.parent != nil {
		return m.parent.Count(filters...)
	}
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, e := range m.entries {
		allOk := true
		for _, f := range filters {
			if !f.check(e) {
				allOk = false
				break
			}
		}
		if allOk {
			count++
		}
	}

	return count
}

func (m *MockLog) Messages(filters ...filter) []string {
	if m.parent != nil {
		return m.parent.Messages(filters...)
	}
	messages := make([]string, 0)

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, e := range m.entries {
		allOk := true
		for _, f := range filters {
			if !f.check(e) {
				allOk = false
				break
			}
		}
		if allOk {
			messages = append(messages, e.Message)
		}
	}

	return messages
}

func (m *MockLog) Message(filters ...filter) (string, bool) {
	if m.parent != nil {
		return m.parent.Message(filters...)
	}
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, e := range m.entries {
		allOk := true
		for _, f := range filters {
			if !f.check(e) {
				allOk = false
				break
			}
		}
		if allOk {
			return e.Message, true
		}
	}

	return "", false
}

// Dump produces a string representation of a mock log, suitable for including
// in assertion/expectation failure messages
func (m *MockLog) Dump() string {
	if m.parent != nil {
		return m.parent.Dump()
	}
	var contents bytes.Buffer

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, e := range m.entries {
		contents.Write(e.toPlainText())
	}
	return contents.String()
}

func (m *MockLog) mergeArgs(args []interface{}) []interface{} {
	if m.fields.contents == nil {
		return args
	}
	mergedArgs := make([]interface{}, 0, len(args)+len(m.fields.contents))
	for _, f := range m.fields.contents {
		mergedArgs = append(mergedArgs, f)
	}
	return append(mergedArgs, args...)
}

func (m *MockLog) addEntry(level Level, prefix string,
	args []interface{}) *Entry {
	if m.parent != nil {
		return m.parent.addEntry(level, prefix, args)
	}
	entry := makeEntry(level, prefix, args)

	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.entries = append(m.entries, entry)
	return entry
}

func (m *MockLog) addFormattedEntry(
	level Level, prefix string, pattern string, args []interface{},
) *Entry {
	if m.parent != nil {
		return m.parent.addFormattedEntry(level, prefix, pattern, args)
	}

	entry := makeFormattedEntry(level, prefix, pattern, args)

	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.entries = append(m.entries, entry)
	return entry
}

// Extend returns a new sub logger by extending the current one with
// extra fields.
func (m *MockLog) Extend(f ...F) Log {
	newFields := Fields{}

	for _, fld := range m.fields.contents {
		newFields.set(fld)
	}

	for _, fld := range f {
		newFields.set(fld)
	}

	ext := &MockLog{
		fields: newFields,
	}
	if m.parent == nil {
		ext.parent = m
	} else {
		ext.parent = m.parent
	}
	return ext
}

// ExtendPrefix returns a new sub logger by extending the current one with
// prefix and extra fields.
func (m *MockLog) ExtendPrefix(prefix string, f ...F) Log {
	ext := m.Extend(f...).(*MockLog)
	ext.prefix += prefix
	return ext
}

// Trace logs a message at trace level
func (m *MockLog) Trace(args ...interface{}) {
	m.addEntry(LevelTrace, m.prefix, m.mergeArgs(args))
}

// Traceln logs a message at trace level
func (m *MockLog) Traceln(args ...interface{}) {
	m.addEntry(LevelTrace, m.prefix, m.mergeArgs(args))
}

// Tracef logs a formatted message at trace level
func (m *MockLog) Tracef(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelTrace, m.prefix, pattern, m.mergeArgs(args))
}

// Debug logs a message at debug level
func (m *MockLog) Debug(args ...interface{}) {
	m.addEntry(LevelDebug, m.prefix, m.mergeArgs(args))
}

// Debugln logs a message at debug level
func (m *MockLog) Debugln(args ...interface{}) {
	m.addEntry(LevelDebug, m.prefix, m.mergeArgs(args))
}

// Debugf logs a formatted message at debug level
func (m *MockLog) Debugf(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelDebug, m.prefix, pattern, m.mergeArgs(args))
}

// Info logs a message at info level
func (m *MockLog) Info(args ...interface{}) {
	m.addEntry(LevelInfo, m.prefix, m.mergeArgs(args))
}

// Infoln logs a message at info level
func (m *MockLog) Infoln(args ...interface{}) {
	m.addEntry(LevelInfo, m.prefix, m.mergeArgs(args))
}

// Infof logs a formatted message at info level
func (m *MockLog) Infof(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelInfo, m.prefix, pattern, m.mergeArgs(args))
}

// Print logs a message at info level
func (m *MockLog) Print(args ...interface{}) {
	m.addEntry(LevelInfo, m.prefix, m.mergeArgs(args))
}

// Println logs a message at info level
func (m *MockLog) Println(args ...interface{}) {
	m.addEntry(LevelInfo, m.prefix, m.mergeArgs(args))
}

// Printf logs a formatted message at info level
func (m *MockLog) Printf(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelInfo, m.prefix, pattern, m.mergeArgs(args))
}

// Warn logs a message at warn level
func (m *MockLog) Warn(args ...interface{}) {
	m.addEntry(LevelWarn, m.prefix, m.mergeArgs(args))
}

// Warnln logs a message at warn level
func (m *MockLog) Warnln(args ...interface{}) {
	m.addEntry(LevelWarn, m.prefix, m.mergeArgs(args))
}

// Warnf logs a formatted message at warn level
func (m *MockLog) Warnf(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelWarn, m.prefix, pattern, m.mergeArgs(args))
}

// Warning logs a message at warn level
func (m *MockLog) Warning(args ...interface{}) {
	m.addEntry(LevelWarn, m.prefix, m.mergeArgs(args))
}

// Warningln logs a message at warn level
func (m *MockLog) Warningln(args ...interface{}) {
	m.addEntry(LevelWarn, m.prefix, m.mergeArgs(args))
}

// Warningf logs a formatted message at warn level
func (m *MockLog) Warningf(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelWarn, m.prefix, pattern, m.mergeArgs(args))
}

// Error logs a message at error level
func (m *MockLog) Error(args ...interface{}) {
	m.addEntry(LevelError, m.prefix, m.mergeArgs(args))
}

// Errorln logs a message at error level
func (m *MockLog) Errorln(args ...interface{}) {
	m.addEntry(LevelError, m.prefix, m.mergeArgs(args))
}

// Errorf logs a formatted message at error level
func (m *MockLog) Errorf(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelError, m.prefix, pattern, m.mergeArgs(args))
}

// Fatal logs a message at fatal level
func (m *MockLog) Fatal(args ...interface{}) {
	panic(m.addEntry(LevelFatal, m.prefix, m.mergeArgs(args)).Message)
}

// Fatalln logs a message at fatal level
func (m *MockLog) Fatalln(args ...interface{}) {
	panic(m.addEntry(LevelFatal, m.prefix, m.mergeArgs(args)).Message)
}

// Fatalf logs a formatted message at fatal level
func (m *MockLog) Fatalf(pattern string, args ...interface{}) {
	panic(m.addFormattedEntry(LevelError, m.prefix, pattern, m.mergeArgs(args)).Message)
}

// Panic logs a message at fatal level and panics
func (m *MockLog) Panic(args ...interface{}) {
	panic(m.addEntry(LevelFatal, m.prefix, m.mergeArgs(args)).Message)
}

// Panicln logs a message at fatal level and panics
func (m *MockLog) Panicln(args ...interface{}) {
	panic(m.addEntry(LevelFatal, m.prefix, m.mergeArgs(args)).Message)
}

// Panicf logs a formatted message at fatal level and panics
func (m *MockLog) Panicf(pattern string, args ...interface{}) {
	panic(m.addFormattedEntry(LevelFatal, m.prefix, pattern, m.mergeArgs(args)).Message)
}
