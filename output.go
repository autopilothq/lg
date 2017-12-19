package lg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"

	multierror "github.com/hashicorp/go-multierror"
)

// hookFn is the handler function for a hook. It returns an error if
// the handler failed.
type hookFn func(*Entry) (err error)

type hook struct {
	fn      hookFn
	options *Options
}

var (
	mutex          sync.RWMutex
	hookFnPointers map[uintptr]hook
	outputs        map[io.Writer]uintptr
)

func makePlainTexthookFn(output io.Writer, options *Options) hookFn {
	return func(e *Entry) (err error) {
		if options.minLevel > e.Level {
			return nil
		}

		var n int
		data := e.toPlainText()
		n, err = output.Write(data)
		if err != nil {
			return err
		}

		if n != len(data) {
			return io.ErrShortWrite
		}

		return nil
	}
}

func makePlainJSONhookFn(output io.Writer, options *Options) hookFn {
	return func(e *Entry) (err error) {
		if options.minLevel > e.Level {
			return nil
		}

		var n int
		data := e.toJSON()
		n, err = output.Write(data)
		if err != nil {
			return
		}

		if n != len(data) {
			return io.ErrShortWrite
		}

		return nil
	}
}

func makeOutputHookFn(output io.Writer, options *Options) hookFn {
	switch options.format {
	case FormatPlainText:
		return makePlainTexthookFn(output, options)

	case FormatJSON:
		return makePlainJSONhookFn(output, options)

	default:
		panic(fmt.Errorf("Invalid log output format %#v", options.format))
	}
}

// AddOutput causes logging to be written to the given io.Writer
func AddOutput(output io.Writer, opts ...func(*Options)) {
	options := makeOptions(opts...)
	mutex.Lock()
	defer mutex.Unlock()

	_, exists := outputs[output]
	if exists {
		panic(errors.New("output is already in use"))
	}

	fn := makeOutputHookFn(output, options)

	outputs[output] = addHook(fn, options)
}

// SetOutput adds or replaces output to the given io.Writer
func SetOutput(output io.Writer, opts ...func(*Options)) {
	options := makeOptions(opts...)
	mutex.Lock()
	defer mutex.Unlock()

	p, exists := outputs[output]
	if exists {
		delete(hookFnPointers, p)
		delete(outputs, output)
	}

	fn := makeOutputHookFn(output, options)

	outputs[output] = addHook(fn, options)

}

// RemoveOutput removes a previously added output
func RemoveOutput(output io.Writer) {
	mutex.Lock()
	defer mutex.Unlock()
	p, exists := outputs[output]
	if exists {
		delete(outputs, output)
		delete(hookFnPointers, p)
	}
}

// RemoveAllOutputs removes all previously added outputs
func RemoveAllOutputs() {
	mutex.Lock()
	defer mutex.Unlock()

	for _, p := range outputs {
		delete(hookFnPointers, p)
	}

	outputs = make(map[io.Writer]uintptr)
}

// AddHook causes logging activity to invoke the given hook function
func AddHook(fn hookFn, opts ...func(*Options)) {
	options := makeOptions(opts...)
	mutex.Lock()
	defer mutex.Unlock()
	addHook(fn, options)
}

func addHook(fn hookFn, options *Options) uintptr {
	p := reflect.ValueOf(fn).Pointer()
	_, exists := hookFnPointers[p]
	if exists {
		panic(errors.New("logging hook already in use"))
	}

	hookFnPointers[p] = hook{fn, options}
	return p
}

// RemoveHook removes a previously added hook function
func RemoveHook(fn hookFn) {
	mutex.Lock()
	defer mutex.Unlock()
	p := reflect.ValueOf(fn).Pointer()
	delete(hookFnPointers, p)
}

func callHooks(entry *Entry) (err error) {
	mutex.RLock()
	defer mutex.RUnlock()
	for _, hook := range hookFnPointers {
		if herr := hook.fn(entry); herr != nil {
			err = multierror.Append(err, herr)
		}
	}

	return err
}

func addEntry(level Level, args []interface{}) (*Entry, error) {
	entry := makeEntry(level, args)
	if err := callHooks(entry); err != nil {
		return nil, err
	}

	return entry, nil
}

func addFormattedEntry(
	level Level, pattern string, args []interface{},
) (*Entry, error) {
	entry := makeFormattedEntry(level, pattern, args)
	if err := callHooks(entry); err != nil {
		return nil, err
	}

	return entry, nil
}

func init() {
	// construct variables for tracking outputs and hooks
	hookFnPointers = make(map[uintptr]hook)
	outputs = make(map[io.Writer]uintptr)

	// set default output
	AddOutput(os.Stdout)
}
