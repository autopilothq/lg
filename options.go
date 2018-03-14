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
