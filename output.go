package lg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"sync/atomic"

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
	mutex   sync.RWMutex
	hookFns map[uint32]hook
	outputs map[io.Writer]uint32
	nextHookID uint32
)

func makePlainTexthookFn(output io.Writer, options *Options) hookFn {
	return func(e *Entry) (err error) {
		if shouldSkip(e, options) {
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
		if shouldSkip(e, options) {
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

func shouldSkip(e *Entry, options *Options) bool {
	if options.minLevels != nil {
		for _, prefixLevel := range options.minLevels {
			if strings.HasPrefix(e.Prefix, prefixLevel.prefix) {
				if prefixLevel.minLevel > e.Level {
					return true
				}
				break
			}
		}
	}
	return false
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

	hookID, exists := outputs[output]
	if exists {
		delete(hookFns, hookID)
		delete(outputs, output)
	}

	fn := makeOutputHookFn(output, options)

	outputs[output] = addHook(fn, options)

}

// RemoveOutput removes a previously added output
func RemoveOutput(output io.Writer) {
	mutex.Lock()
	defer mutex.Unlock()
	hookID, exists := outputs[output]
	if exists {
		delete(outputs, output)
		delete(hookFns, hookID)
	}
}

// RemoveAllOutputs removes all previously added outputs
func RemoveAllOutputs() {
	mutex.Lock()
	defer mutex.Unlock()

	for _, p := range outputs {
		delete(hookFns, p)
	}

	outputs = make(map[io.Writer]uint32)
}

// AddHook causes logging activity to invoke the given hook function.
// It returns an id which can be used to remove the hook with RemoveHook.
func AddHook(fn hookFn, opts ...func(*Options)) uint32 {
	options := makeOptions(opts...)
	mutex.Lock()
	defer mutex.Unlock()
	return addHook(fn, options)
}

func addHook(fn hookFn, options *Options) uint32 {
	hookID := atomic.AddUint32(&nextHookID, uint32(1))
	hookFns[hookID] = hook{fn, options}
	return hookID
}

// RemoveHook removes a previously added hook function
func RemoveHook(hookID uint32) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(hookFns, hookID)
}

func callHooks(entry *Entry) (err error) {
	mutex.RLock()
	defer mutex.RUnlock()
	for _, hook := range hookFns {
		if herr := hook.fn(entry); herr != nil {
			err = multierror.Append(err, herr)
		}
	}

	return err
}

func addEntry(level Level, prefix string, args []interface{}) (*Entry, error) {
	entry := makeEntry(level, prefix, args)
	if err := callHooks(entry); err != nil {
		return nil, err
	}

	return entry, nil
}

func addFormattedEntry(
	level Level, prefix string, pattern string, args []interface{},
) (*Entry, error) {
	entry := makeFormattedEntry(level, prefix, pattern, args)
	if err := callHooks(entry); err != nil {
		return nil, err
	}

	return entry, nil
}

func init() {
	// construct variables for tracking outputs and hooks
	hookFns = make(map[uint32]hook)
	outputs = make(map[io.Writer]uint32)

	// set default output
	AddOutput(os.Stdout)
}
