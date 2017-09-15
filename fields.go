package lg

import (
	"encoding/json"
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
		out += fld.Key + ":" + renderMessage(fld.Val)
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

// MarshalJSON allows Fields to satisfy json.Marshalable
func (f *Fields) MarshalJSON() ([]byte, error) {
	if len(f.contents) == 0 {
		return []byte("{}"), nil
	}

	out := ""
	for _, fld := range f.contents {
		if out != "" {
			out += ","
		}
		keyJSON, err := json.Marshal(fld.Key)
		if err != nil {
			return []byte(nil), err
		}
		valueJSON, err := json.Marshal(fld.Val)
		if err != nil {
			return []byte(nil), err
		}
		out += string(keyJSON) + ":" + string(valueJSON)
	}
	return []byte("{" + out + "}"), nil
}
