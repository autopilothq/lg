package lg_test

import (
	"regexp"

	"github.com/autopilothq/lg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("log mocking", func() {
	var (
		log     lg.Log
		mockLog *lg.MockLog
	)

	BeforeEach(func() {
		mockLog = lg.Mock()
		log = mockLog
	})

	Context("Count()", func() {

		It("counts log messages", func() {
			Expect(mockLog.Count()).To(Equal(0))
			log.Info("banana")
			Expect(mockLog.Count()).To(Equal(1))
		})

		It("counts log messages, filtered by exact level", func() {
			log.Info("banana")
			log.Warn("starfruit")
			Expect(mockLog.Count(lg.AtLevel(lg.LevelWarn))).To(Equal(1))
			Expect(mockLog.Count(lg.AtLevel(lg.LevelInfo))).To(Equal(1))
			Expect(mockLog.Count(lg.AtLevel(lg.LevelFatal))).To(Equal(0))
		})

		It("counts log messages, filtered by minimum level", func() {
			log.Info("banana")
			log.Warn("starfruit")
			log.Error("pineapple")
			Expect(mockLog.Count(lg.AtLeastLevel(lg.LevelInfo))).To(Equal(3))
			Expect(mockLog.Count(lg.AtLeastLevel(lg.LevelWarn))).To(Equal(2))
			Expect(mockLog.Count(lg.AtLeastLevel(lg.LevelError))).To(Equal(1))
			Expect(mockLog.Count(lg.AtLeastLevel(lg.LevelFatal))).To(Equal(0))
		})

		It("counts log messages, filtered by regexp", func() {
			log.Infof("a number: %d", 1234567)
			log.Warnf("also a number: %d", 1234567)
			log.Debug("just some text")
			Expect(mockLog.Count(lg.Regexp("12345"))).To(Equal(2))
		})

		It("counts log messages, filtered by contents", func() {
			log.Trace("foobar")
			log.Error("barbaz")
			log.Info("foobaz")
			Expect(mockLog.Count(lg.Contains("bar"))).To(Equal(2))
		})

		It("counts log messages, with multiple filters", func() {
			log.Trace("foobar")
			log.Error("barbaz")
			log.Warning("foobaz")
			Expect(
				mockLog.Count(lg.AtLevel(lg.LevelTrace), lg.Contains("bar")),
			).To(Equal(1))
		})

	})

	Context("Message()", func() {

		It("fetches the first matching message contents", func() {
			log.Error("uh oh")
			message, found := mockLog.Message()
			Expect(found).To(Equal(true))
			Expect(message).To(Equal("uh oh"))
		})

		It("fetches the first matching message contents, filtered by exact level",
			func() {
				log.Print("blah blah blah")
				log.Error("uh oh")
				log.Print("blah blah blah")
				message, found := mockLog.Message(lg.AtLevel(lg.LevelError))
				Expect(found).To(Equal(true))
				Expect(message).To(Equal("uh oh"))
			})

		It("fetches the first matching message contents, filtered by level",
			func() {
				log.Print("blah blah blah")
				log.Warn("eep")
				log.Error("uh oh")
				message, found := mockLog.Message(lg.AtLeastLevel(lg.LevelWarn))
				Expect(found).To(Equal(true))
				Expect(message).To(Equal("eep"))
			})

	})

	Context("Messages()", func() {

		It("fetches the first matching message contents", func() {
			log.Error("uh oh")
			message, found := mockLog.Message()
			Expect(found).To(Equal(true))
			Expect(message).To(Equal("uh oh"))
		})

		It("fetches the first matching message contents, filtered by exact level",
			func() {
				log.Print("blah blah blah")
				log.Error("uh oh")
				log.Print("blah blah blah")
				message, found := mockLog.Message(lg.AtLevel(lg.LevelError))
				Expect(found).To(Equal(true))
				Expect(message).To(Equal("uh oh"))
			})

		It("fetches the first matching message contents, filtered by level",
			func() {
				log.Print("blah blah blah")
				log.Warn("eep")
				log.Error("uh oh")
				message, found := mockLog.Message(lg.AtLeastLevel(lg.LevelWarn))
				Expect(found).To(Equal(true))
				Expect(message).To(Equal("eep"))
			})

	})

	Context("Dump()", func() {

		It("returns the plain text representation of the captured log entries",
			func() {
				log.Print("words")
				log.Debug("things happening", lg.F{"foo", "bar"})

				expected := "info  words\ndebug [foo:bar] things happening\n"

				timePattern := regexp.MustCompile("[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\\.[0-9]{3} ")

				// Trim times out of dumped logs
				original := mockLog.Dump()
				actual := timePattern.ReplaceAllLiteralString(original, "")

				Expect(actual).To(Equal(expected))
			})

	})

})
