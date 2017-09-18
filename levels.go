package lg

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Level is a log level enumerable
type Level uint

// Log level literals
const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	default:
		return ""
	}
}

func (l Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

// ParseLevel transforms a string log level into a Level type. Returns an error
// if the given level is invalid.
func ParseLevel(level string) (Level, error) {
	switch strings.ToLower(level) {
	case "trace":
		return LevelTrace, nil
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	case "fatal":
		return LevelFatal, nil
	default:
		return Level(0), fmt.Errorf("Invalid log level '%s'", level)
	}
}
