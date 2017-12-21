package lg

import (
	"bytes"

	"github.com/autopilothq/lg/encoding"
	fancy "github.com/autopilothq/lg/encoding/json"
	text "github.com/autopilothq/lg/encoding/text"
)

// F represents a single pair of log 'fields'
type F struct {
	Key string
	Val interface{}
}

// Fields is a slice of field key-value pairs
type Fields struct {
	contents []F
}

// Len returns the number of fields
func (f *Fields) Len() int {
	return len(f.contents)
}

func (f *Fields) renderPlainText() string {
	if len(f.contents) == 0 {
		return ""
	}

	out := bytes.NewBufferString("[")
	// out := ""

	for i, fld := range f.contents {

		if i > 0 {
			// out += " "
			out.WriteByte(' ')
		}

		out.WriteString(fld.Key)
		out.WriteByte(':')
		out.WriteString(RenderMessage(fld.Val))
		// out += fld.Key + ":" + RenderMessage(fld.Val)
	}

	// return "[" + out + "] "
	out.WriteString("] ")
	return out.String()
}

func (f *Fields) set(fld F) {
	for idx := range f.contents {
		if f.contents[idx].Key == fld.Key {
			f.contents[idx].Val = fld.Val
			return
		}
	}
	f.contents = append(f.contents, fld)
}

// encodeJSON allows Fields to be marshaled to JSON via the encoder
func (f *Fields) encodeJSON(enc *fancy.Encoder) (err error) {
	if len(f.contents) == 0 {
		return nil
	}

	if err = enc.StartObject(); err != nil {
		return err
	}

	for _, fld := range f.contents {
		err := encoding.EncodeKeyValue(enc, fld.Key, fld.Val)
		if err != nil {
			return err
		}
	}

	if err = enc.EndObject(); err != nil {
		return err
	}

	return nil
}

// encodeText allows Fields to be marshaled to TEXT via the encoder
func (f *Fields) encodeText(enc *text.Encoder) (err error) {
	if len(f.contents) == 0 {
		return nil
	}

	for _, fld := range f.contents {
		err := encoding.EncodeKeyValue(enc, fld.Key, fld.Val)
		if err != nil {
			return err
		}
	}

	return nil
}
