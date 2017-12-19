package lg

import (
	"bytes"

	"github.com/autopilothq/lg/encoding"
	fancy "github.com/autopilothq/lg/encoding/json"
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
	if f.contents != nil {
		for _, item := range f.contents {
			if item.Key == fld.Key {
				item.Val = fld.Val
				return
			}
		}
	} else {
		f.contents = make([]F, 0)
	}

	f.contents = append(f.contents, fld)
}

// encodeJSON allows Fields to be marshaled to JSON via the encoder
func (f *Fields) encodeJSON(encoder *fancy.Encoder) (err error) {
	if len(f.contents) > 0 {
		for _, fld := range f.contents {
			err := encoding.EncodeKeyValue(encoder, fld.Key, fld.Val)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
