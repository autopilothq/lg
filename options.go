package lg

import (
	"fmt"
	"regexp"
)

type OutputFormat uint

type PrefixLevel struct {
	prefix   string
	minLevel Level
}

type Options struct {
	minLevels []PrefixLevel
	format    OutputFormat
}

const (
	FormatPlainText OutputFormat = iota
	FormatJSON
)

var (
	prefixLevelPattern = regexp.MustCompile("(?:([^\\s,()]+)=)?(trace|debug|info|warn|error|fatal)")
)

func makeOptions(opts ...func(*Options)) *Options {
	result := &Options{
		format:    FormatPlainText,
		minLevels: nil,
	}
	for _, opt := range opts {
		opt(result)
	}
	return result
}

func PlainText() func(o *Options) {
	return func(o *Options) {
		o.format = FormatPlainText
	}
}

func JSON() func(*Options) {
	return func(o *Options) {
		o.format = FormatJSON
	}
}

func MinLevel(l Level) func(*Options) {
	return func(o *Options) {
		if o.minLevels == nil {
			o.minLevels = make([]PrefixLevel, 1)
			o.minLevels[0] = PrefixLevel{prefix: "", minLevel: l}
		} else if len(o.minLevels) == 0 {
			o.minLevels = append(o.minLevels, PrefixLevel{prefix: "", minLevel: l})
		} else {
			last := o.minLevels[len(o.minLevels)-1]
			if last.prefix == "" {
				last.minLevel = l
			} else {
				o.minLevels = append(o.minLevels, PrefixLevel{prefix: "", minLevel: l})
			}
		}
	}
}

// Levels specifies the minimum log level for an Output. Levels can be specfied by log prefix.
//
// Examples:
//
//   // Show debug level and above on stdout
//   lg.SetOutput(os.Stdout, lg.Levels("debug"))
//
//   // Show info level and above by default, but only show warn and above
//   // for any logs from the "Server" prefix, including sub-prefixes.
//   lg.SetOutput(os.Stdout, lg.Levels("(Server=warn) info"))
//
//   // Show errors and above, except for logs with the "Request" prefix,
//   // which will show trace and above
//   lg.SetOutput(os.Stdout, lg.Levels("(Request=trace) error"))
//
// The levels string is evaluated left-to-right, so more specific prefixes
// should be to the left of more general ones.
func Levels(levels string) func(*Options) {
	return func(o *Options) {
		matches := prefixLevelPattern.FindAllStringSubmatch(levels, -1)
		o.minLevels = make([]PrefixLevel, len(matches))
		for n, match := range matches {
			l, err := ParseLevel(match[2])
			if err != nil {
				panic(fmt.Errorf("Unparsable levels string '%s': '%s' is not a valid level", levels, match[2]))
			}
			o.minLevels[n] = PrefixLevel{prefix: match[1], minLevel: l}
		}
	}
}
