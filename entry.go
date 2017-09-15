package lg

import (
	"encoding/json"
	"fmt"
	"time"
)

// Entry represents a log entry
type Entry struct {
	Timestamp time.Time `json:"t"`
	Message   string    `json:"m"`
	Level     Level     `json:"l,string"`
	Fields    *Fields   `json:"f"`
}

const (
	InvalidTime = "<invalid time>"
)

func (e *Entry) toPlainText() []byte {
	timeBytes, err := e.Timestamp.MarshalText()
	if err != nil {
		timeBytes = []byte(InvalidTime)
	}
	return append(append(timeBytes, ' '), e.toPlainTextWithoutTime()...)
}

func (e *Entry) toPlainTextWithoutTime() []byte {
	return []byte((e.Level.String() + "      ")[0:5] + " " +
		e.Fields.renderPlainText() +
		e.Message + "\n")
}

func (e *Entry) toJSON() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		return []byte("{\"error\":\"encoding error\"}\n")
	}
	return append(bytes, '\n')
}

func makeEntry(level Level, args []interface{}) *Entry {
	fields, remaining := extractAllFields(args)

	message := renderMessage(remaining...)

	return &Entry{
		time.Now().UTC(),
		message,
		level,
		fields,
	}
}

func makeFormattedEntry(
	level Level, pattern string, args []interface{},
) *Entry {
	fields, remaining := extractTrailingFields(args)

	message := fmt.Sprintf(pattern, remaining...)

	return &Entry{
		time.Now().UTC(),
		message,
		level,
		fields,
	}
}
