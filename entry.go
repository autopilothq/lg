package lg

import (
	"fmt"
	"time"

	fancy "github.com/autopilothq/lg/encoding/json"
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
	encoder := fancy.NewEncoder()
	err := e.Fields.MarshalJSONEncoder(encoder)
	if err != nil {
		return []byte("{\"error\":\"encoding error\"}\n")
	}
	return encoder.Bytes()
}

func makeEntry(level Level, args []interface{}) *Entry {
	fields, remaining := ExtractAllFields(args)

	message := RenderMessage(remaining...)

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
	fields, remaining := ExtractTrailingFields(args)

	message := fmt.Sprintf(pattern, remaining...)

	return &Entry{
		time.Now().UTC(),
		message,
		level,
		fields,
	}
}
