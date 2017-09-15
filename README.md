## Lg

A simple, structured logger for Go.



### Basic Usage

Lg provides timestamped logging with levels and/or Printf formatting:

```go
package main

import (
  "github.com/autopilothq/lg"
)

func main() {
  
  // Simple logging
  lg.Println("starting")
  
  // Levels (trace, debug, info, warn, error, fatal)
  lg.Warn("danger")
  
  // Printf formatting
  thing := &map[string]string{
    "foo": "bar",
  }
  lg.Debugf("what is this %#v", thing)
  
}
```

Will output:

```
2017-09-15T00:08:54.851998345Z info  starting
2017-09-15T00:08:54.852009867Z warn  danger
2017-09-15T00:08:54.852022087Z debug what is this &map[string]string{"foo":"bar"}
```



### Fields

Fields (key-value pairs) can be added to logging to provide context:

```go
  // Add fields for context
  lg.Debug("things happened",
           lg.F{"traceid", "1234abcd"},
           lg.F{"module", "thinginator"})

```

Will output:

```
2017-09-15T00:16:43.848278652Z debug [traceid:1234abcd module:thinginator] things happened
```





### Extending

Custom loggers can be created with pre-defined fields:

```go
  // Extend logs to add default fields

  log1 := lg.Extend(lg.F{"traceid", "1234abcd"})

  log2 := log1.Extend(lg.F{"module", "thinginator"})

  log2.Debug("things happened")
```

Will output:

```
2017-09-15T00:19:16.813565094Z debug [traceid:1234abcd module:thinginator] things happened
```



Custom loggers conform to the interface `lg.Log` which makes them easy to store and re-use.





### Output



By default, lg will output in plain text format to stdout. You can add and remove outputs as needed, configure their log levels and/or use JSON format:



```go
// remove default stdout output
lg.RemoveOutput(os.Stdout) 

// write logging at level info or higher to a file
f, err := os.OpenFile("/tmp/log", os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
  panic(err)
}
lg.AddOutput(f, lg.MinLevel(lg.LevelInfo))

// write logging at level error or higher to stderr
lg.AddOutput(os.Stderr, lg.MinLevel(lg.LevelError))

// write logging at any level to stdout, in JSON format
lg.AddOutput(os.Stdout, lg.JSON())
```





### Hooks

Hooks are similar to outputs, but instead of writing to an output stream, a hook function is called with a log entry.



```go
handler = func(e *lg.Entry) {
  fmt.Println(e.Timestamp, e.Level.String(), e.Message, e.Fields)
}

lg.AddHook(handler)

// A hook can be removed by passing the same handler function to lg.RemoveHook:
// lg.RemoveHook(handler)
```





### Mocking

Lg's strangest feature is it's mock loggers. But they can be very helpful for testing.

Say you have a function you want to unit test:

```go
func ChurnSomeNumbers(log lg.Log, a int, b int) (int, error) {
  if a < b {
    err := fmt.Errorf("%d cannot be less than %d", a, b)
    log.Error(err)
    return 0, err
  }
  
  return a - b, nil
}
```



As well as testing passing cases, you'll also want to test failing cases, but that means your test output will contain errors, which are expected. Using a mock log can help here, you can silence output when the test passes, verify the function under test logs the way you expect it to, and you can still see the log output on failure, by including it in an assertion/expectation failure message:

```go
Describe("ChurnSomeNumbers()", func() {
  BeforeEach({
    log := lg.Mock()
  })
  
  It("fails when a is less than b", func() {
    _, err := ChurnSomeNumbers(log, 3, 5)
    Expect(err).To(HaveOccurred())
    Expect(log.Count(
      lg.AtLevel(lg.LevelError),
    )).To(Equal(1), "log should contain exactly one error:\n%s", log.Dump())
  })
})
```

If the test passes, it will produce no extra output, but if it fails, the log will be included in the expectation failure message.



### Running tests

To run tests, you'll need ginkgo and gomega:

```
$ go get github.com/onsi/ginkgo github.com/onsi/gomega
```

Then run the tests with ginkgo:

```
$ ginkgo
```
