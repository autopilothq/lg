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
	if f.contents == nil || len(f.contents) == 0 {
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

func (f *Fields) MarshalJSON() ([]byte, error) {
	if f.contents == nil || len(f.contents) == 0 {
		return []byte("{}"), nil
	}

	out := ""
	for _, fld := range f.contents {
		if out != "" {
			out += ","
		}
		keyJson, err := json.Marshal(fld.Key)
		if err != nil {
			return []byte(nil), err
		}
		valueJson, err := json.Marshal(fld.Val)
		if err != nil {
			return []byte(nil), err
		}
		out += string(keyJson) + ":" + string(valueJson)
	}
	return []byte("{" + out + "}"), nil
}
