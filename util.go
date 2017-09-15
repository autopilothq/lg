package lg

import (
	"fmt"
)

func extractAllFields(
	args []interface{},
) (fields *Fields, remaining []interface{}) {
	fields = &Fields{}
	remaining = make([]interface{}, 0)

	for _, a := range args {
		switch v := a.(type) {
		case F:
			fields.set(v)

		default:
			remaining = append(remaining, a)
		}
	}

	return fields, remaining
}

func extractTrailingFields(
	args []interface{},
) (fields *Fields, remaining []interface{}) {

	extractedFs := make([]F, 0)
	fields = &Fields{}
	remaining = make([]interface{}, 0)

	for i := len(args) - 1; i >= 0; i-- {
		a := args[i]
		switch v := a.(type) {
		case F:
			extractedFs = append([]F{v}, extractedFs...)

		default:
			remaining = append(remaining, a)
		}
	}

	for _, ef := range extractedFs {
		fields.set(ef)
	}

	return fields, remaining
}

func renderMessage(args ...interface{}) string {
	message := ""
	for _, a := range args {
		if message != "" {
			message += " "
		}
		s, ok := a.(fmt.Stringer)
		if ok {
			message += s.String()
		} else {
			message += fmt.Sprintf("%v", a)
		}
	}
	return message
}
