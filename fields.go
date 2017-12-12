package lg

import (
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
	out := ""
	for _, fld := range f.contents {
		if out != "" {
			out += " "
		}
		out += fld.Key + ":" + RenderMessage(fld.Val)
	}
	return "[" + out + "] "
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
func (f *Fields) encodeJSON(encoder *fancy.Encoder) error {
	if len(f.contents) == 0 {
		return encoder.AddByteString("{}")
	}

	for _, fld := range f.contents {
		err := fancy.EncodeKeyValue(encoder, fld.Key, fld.Val)
		if err != nil {
			return err
		}
	}

	return nil
}
