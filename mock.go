package lg

import (
	"bytes"
	"sync"
)

type MockLog struct {
	entries []*Entry
	mutex   sync.RWMutex
}

type filter interface {
	check(*Entry) bool
}

func Mock() *MockLog {
	return &MockLog{}
}

func (m *MockLog) Count(filters ...filter) (count int) {
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
	contents := bytes.NewBuffer([]byte{})

	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, e := range m.entries {
		contents.Write(e.toPlainText())
	}
	return contents.String()
}

func (m *MockLog) addEntry(level Level, args []interface{}) *Entry {
	entry := makeEntry(level, args)

	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.entries = append(m.entries, entry)
	return entry
}

func (m *MockLog) addFormattedEntry(
	level Level, pattern string, args []interface{},
) *Entry {
	entry := makeFormattedEntry(level, pattern, args)

	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.entries = append(m.entries, entry)
	return entry
}

func (m *MockLog) Extend(f ...F) Log {
	// TODO: incorporate fields properly
	return m
}

// Trace logs a message at trace level
func (m *MockLog) Trace(args ...interface{}) {
	m.addEntry(LevelTrace, args)
}

// Traceln logs a message at trace level
func (m *MockLog) Traceln(args ...interface{}) {
	m.addEntry(LevelTrace, args)
}

// Tracef logs a formatted message at trace level
func (m *MockLog) Tracef(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelTrace, pattern, args)
}

// Debug logs a message at debug level
func (m *MockLog) Debug(args ...interface{}) {
	m.addEntry(LevelDebug, args)
}

// Debugln logs a message at debug level
func (m *MockLog) Debugln(args ...interface{}) {
	m.addEntry(LevelDebug, args)
}

// Debugf logs a formatted message at debug level
func (m *MockLog) Debugf(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelDebug, pattern, args)
}

// Info logs a message at info level
func (m *MockLog) Info(args ...interface{}) {
	m.addEntry(LevelInfo, args)
}

// Infoln logs a message at info level
func (m *MockLog) Infoln(args ...interface{}) {
	m.addEntry(LevelInfo, args)
}

// Infof logs a formatted message at info level
func (m *MockLog) Infof(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelInfo, pattern, args)
}

// Print logs a message at info level
func (m *MockLog) Print(args ...interface{}) {
	m.addEntry(LevelInfo, args)
}

// Println logs a message at info level
func (m *MockLog) Println(args ...interface{}) {
	m.addEntry(LevelInfo, args)
}

// Printf logs a formatted message at info level
func (m *MockLog) Printf(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelInfo, pattern, args)
}

// Warn logs a message at warn level
func (m *MockLog) Warn(args ...interface{}) {
	m.addEntry(LevelWarn, args)
}

// Warnln logs a message at warn level
func (m *MockLog) Warnln(args ...interface{}) {
	m.addEntry(LevelWarn, args)
}

// Warnf logs a formatted message at warn level
func (m *MockLog) Warnf(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelWarn, pattern, args)
}

// Warning logs a message at warn level
func (m *MockLog) Warning(args ...interface{}) {
	m.addEntry(LevelWarn, args)
}

// Warningln logs a message at warn level
func (m *MockLog) Warningln(args ...interface{}) {
	m.addEntry(LevelWarn, args)
}

// Warningf logs a formatted message at warn level
func (m *MockLog) Warningf(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelWarn, pattern, args)
}

// Error logs a message at error level
func (m *MockLog) Error(args ...interface{}) {
	m.addEntry(LevelError, args)
}

// Errorln logs a message at error level
func (m *MockLog) Errorln(args ...interface{}) {
	m.addEntry(LevelError, args)
}

// Errorf logs a formatted message at error level
func (m *MockLog) Errorf(pattern string, args ...interface{}) {
	m.addFormattedEntry(LevelError, pattern, args)
}

// Fatal logs a message at fatal level
func (m *MockLog) Fatal(args ...interface{}) {
	panic(m.addEntry(LevelFatal, args).Message)
}

// Fatalln logs a message at fatal level
func (m *MockLog) Fatalln(args ...interface{}) {
	panic(m.addEntry(LevelFatal, args).Message)
}

// Fatalf logs a formatted message at fatal level
func (m *MockLog) Fatalf(pattern string, args ...interface{}) {
	panic(m.addFormattedEntry(LevelError, pattern, args).Message)
}

// Panic logs a message at fatal level and panics
func (m *MockLog) Panic(args ...interface{}) {
	panic(m.addEntry(LevelFatal, args).Message)
}

// Panicln logs a message at fatal level and panics
func (m *MockLog) Panicln(args ...interface{}) {
	panic(m.addEntry(LevelFatal, args).Message)
}

// Panicf logs a formatted message at fatal level and panics
func (m *MockLog) Panicf(pattern string, args ...interface{}) {
	panic(m.addFormattedEntry(LevelFatal, pattern, args).Message)
}
