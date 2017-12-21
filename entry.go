package lg

import (
	"bytes"
	"fmt"
	"time"

	json "github.com/autopilothq/lg/encoding"
	fancy "github.com/autopilothq/lg/encoding/json"
	text "github.com/autopilothq/lg/encoding/text"
)

// Entry represents a log entry
type Entry struct {
	Timestamp time.Time `json:"t"`
	Prefix    string    `json:"p,omitempty"`
	Message   string    `json:"m"`
	Level     Level     `json:"l,string"`
	Fields    Fields    `json:"f"`
}

const (
	// TimeFormat is the layout to use when rendering time
	TimeFormat = "2006-01-02T15:04:05.000"
)

func (e *Entry) toPlainText() []byte {
	timeBytes := bytes.NewBufferString(e.Timestamp.Format(TimeFormat))

	_, err := timeBytes.WriteString(e.Level.AlignedString())
	if err != nil {
		return append([]byte(err.Error()), '\n')
	}

	if e.Prefix != "" {
		_, err = timeBytes.WriteString("@" + e.Prefix + " ")
		if err != nil {
			return append([]byte(err.Error()), '\n')
		}
	}

	if e.Fields.Len() > 0 {
		enc := text.NewEncoder()

		if err = enc.StartArray(); err != nil {
			return append([]byte(err.Error()), '\n')
		}

		if err = e.Fields.encodeText(enc); err != nil {
			return append([]byte(err.Error()), '\n')
		}

		if err = enc.EndArray(); err != nil {
			return append([]byte(err.Error()), '\n')
		}

		if _, err = timeBytes.Write(enc.Bytes()); err != nil {
			return append([]byte(err.Error()), '\n')
		}

		if err = timeBytes.WriteByte(' '); err != nil {
			return append([]byte(err.Error()), '\n')
		}
	}

	if _, err = timeBytes.WriteString(e.Message); err != nil {
		return append([]byte(err.Error()), '\n')
	}

	return append(timeBytes.Bytes(), '\n')
}

func makeJSONError(enc *fancy.Encoder, err error) []byte {
	b := fmt.Sprintf(`{"error": "encoding error: %s"}`,
		json.EncodeValue(enc, err))
	return append([]byte(b), '\n')
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

	if e.Prefix != "" {
		err = json.EncodeStringKeyValue(enc, "p", e.Prefix)
		if err != nil {
			return makeJSONError(enc, err)
		}
	}

	if e.Fields.Len() > 0 {
		if err = enc.AddKey("f"); err != nil {
			return makeJSONError(enc, err)
		}

		err = e.Fields.encodeJSON(enc)
		if err != nil {
			return makeJSONError(enc, err)
		}
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

func makeEntry(level Level, prefix string, args []interface{}) *Entry {
	fields, remaining := ExtractAllFields(args)

	message := RenderMessage(remaining...)

	return &Entry{
		Timestamp: time.Now().UTC(),
		Prefix:    prefix,
		Message:   message,
		Level:     level,
		Fields:    fields,
	}
}

func makeFormattedEntry(
	level Level, prefix string, pattern string, args []interface{},
) *Entry {
	fields, remaining := ExtractTrailingFields(args)

	message := fmt.Sprintf(pattern, remaining...)

	return &Entry{
		Timestamp: time.Now().UTC(),
		Prefix:    prefix,
		Message:   message,
		Level:     level,
		Fields:    fields,
	}
}
