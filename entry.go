package lg

import (
	"bytes"
	"fmt"
	"time"

	json "github.com/autopilothq/lg/encoding"
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
	// TimeFormat is the layout to use when rendering time
	TimeFormat = "2006-01-02T15:04:05.999"
)

func (e *Entry) toPlainText() []byte {
	timeBytes := bytes.NewBufferString(e.Timestamp.Format(TimeFormat)).Bytes()
	return append(
		append(timeBytes, ' '),
		e.toPlainTextWithoutTime()...)
}

func (e *Entry) toPlainTextWithoutTime() []byte {
	return []byte((e.Level.String() + "      ")[0:5] + " " +
		e.Fields.renderPlainText() +
		e.Message + "\n")
}

func makeJSONError(enc *fancy.Encoder, err error) []byte {
	b := fmt.Sprintf(`{"error": "encoding error: %s"}`,
		json.EncodeValue(enc, err))
	return []byte(b)
}

func (e *Entry) toJSON() []byte {
	enc := fancy.NewEncoder()

	err := enc.StartObject()
	if err != nil {
		return makeJSONError(enc, err)
	}

	err = json.EncodeTimeKeyValue(enc, "t", e.Timestamp)
	if err != nil {
		return makeJSONError(enc, err)
	}

	err = json.EncodeStringKeyValue(enc, "l", e.Level.String())
	if err != nil {
		return makeJSONError(enc, err)
	}

	err = e.Fields.encodeJSON(enc)
	if err != nil {
		return makeJSONError(enc, err)
	}

	err = json.EncodeStringKeyValue(enc, "m", e.Message)
	if err != nil {
		return makeJSONError(enc, err)
	}

	err = enc.EndObject()
	if err != nil {
		return makeJSONError(enc, err)
	}

	return append(enc.Bytes(), '\n')
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
