package lg

// Log describes the lg logging interface
type Log interface {
	Trace(args ...interface{})
	Traceln(args ...interface{})
	Tracef(pattern string, args ...interface{})

	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(pattern string, args ...interface{})

	Print(args ...interface{})
	Println(args ...interface{})
	Printf(pattern string, args ...interface{})

	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(pattern string, args ...interface{})

	Warn(args ...interface{})
	Warnln(args ...interface{})
	Warnf(pattern string, args ...interface{})

	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(pattern string, args ...interface{})

	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(pattern string, args ...interface{})

	Panic(args ...interface{})
	Panicln(args ...interface{})
	Panicf(pattern string, args ...interface{})

	Extend(f ...F) Log
}
