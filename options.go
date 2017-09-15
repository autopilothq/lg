package lg

type OutputFormat uint

const (
	FormatPlainText OutputFormat = iota
	FormatJSON
)

type Options struct {
	minLevel Level
	format   OutputFormat
}

func makeOptions(opts ...func(*Options)) *Options {
	result := &Options{
		format:   FormatPlainText,
		minLevel: LevelTrace,
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
		o.minLevel = l
	}
}
