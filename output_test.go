package lg_test

import (
	"bytes"
	"os"
	"regexp"
	"strings"

	"github.com/autopilothq/lg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestLogOutput struct {
	bytes.Buffer
}

var (
	lastEntryPattern = regexp.MustCompile("\\S+\n$")
)

func (tlo *TestLogOutput) lastEntry() string {
	return strings.TrimSpace(lastEntryPattern.FindString(tlo.String()))
}

var _ = Describe("log output", func() {

	var tlo *TestLogOutput

	BeforeEach(func() {
		tlo = &TestLogOutput{}
		lg.RemoveOutput(os.Stdout)
		lg.AddOutput(tlo, lg.Levels("Foo.Bar=info (Bar=error)(Foo=error) (trace)"))
	})

	AfterEach(func() {
		lg.RemoveOutput(tlo)
		lg.AddOutput(os.Stdout)
	})

	It("should be filtered by prefix and level", func() {
		fooLog := lg.ExtendWithPrefix("Foo")
		foobarLog := fooLog.ExtendPrefix("Bar")
		barLog := lg.ExtendWithPrefix("Bar")
		lg.Debug("1")
		Expect(tlo.lastEntry()).To(Equal("1"))
		lg.Warn("2")
		Expect(tlo.lastEntry()).To(Equal("2"))
		fooLog.Error("3")
		Expect(tlo.lastEntry()).To(Equal("3"))
		fooLog.Trace("4")
		Expect(tlo.lastEntry()).To(Equal("3"))
		foobarLog.Info("5")
		Expect(tlo.lastEntry()).To(Equal("5"))
		foobarLog.Debug("6")
		Expect(tlo.lastEntry()).To(Equal("5"))
		barLog.Trace("7")
		Expect(tlo.lastEntry()).To(Equal("5"))
		barLog.Warn("8")
		Expect(tlo.lastEntry()).To(Equal("5"))
	})
})
