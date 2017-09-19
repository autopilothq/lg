package lg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"
)

type hook struct {
	fn      func(*Entry)
	options *Options
}

var (
	mutex          sync.RWMutex
	hookFnPointers map[uintptr]hook
	outputs        map[io.Writer]uintptr
)

func makeOutputHookFn(output io.Writer, options *Options) func(*Entry) {
	return func(e *Entry) {
		if options.minLevel > e.Level {
			return
		}

		switch options.format {

		case FormatPlainText:
			output.Write(e.toPlainText())

		case FormatJSON:
			output.Write(e.toJSON())

		default:
			panic(fmt.Errorf("Invalid log output format %#v", options.format))
		}

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

// AddHook causes logging activity to invoke the given hook function
func AddHook(fn func(*Entry), opts ...func(*Options)) {
	options := makeOptions(opts...)
	mutex.Lock()
	defer mutex.Unlock()
	addHook(fn, options)
}

func addHook(fn func(*Entry), options *Options) uintptr {
	p := reflect.ValueOf(fn).Pointer()
	_, exists := hookFnPointers[p]
	if exists {
		panic(errors.New("logging hook already in use"))
	}

	hookFnPointers[p] = hook{fn, options}
	return p
}

// RemoveHook removes a previously added hook function
func RemoveHook(fn func(*Entry)) {
	mutex.Lock()
	defer mutex.Unlock()
	p := reflect.ValueOf(fn).Pointer()
	delete(hookFnPointers, p)
}

func callHooks(entry *Entry) {
	mutex.RLock()
	defer mutex.RUnlock()
	for _, hook := range hookFnPointers {
		hook.fn(entry)
	}
}

func addEntry(level Level, args []interface{}) *Entry {
	entry := makeEntry(level, args)
	callHooks(entry)
	return entry
}

func addFormattedEntry(level Level, pattern string, args []interface{}) *Entry {
	entry := makeFormattedEntry(level, pattern, args)
	callHooks(entry)
	return entry
}

func init() {
	// construct variables for tracking outputs and hooks
	hookFnPointers = make(map[uintptr]hook)
	outputs = make(map[io.Writer]uintptr)

	// set default output
	AddOutput(os.Stdout)
}
