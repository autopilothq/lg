package benchmark

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/autopilothq/lg"
)

func BenchmarkFieldLogging(b *testing.B) {
	var out bytes.Buffer
	lg.RemoveAllOutputs()
	lg.AddOutput(&out, lg.JSON())

	var (
		now      = time.Now()
		before   = now.Add(-30 * time.Minute)
		later    = now.Add(2 * time.Hour)
		err      = errors.New("It broke")
		short    = 100 * time.Millisecond
		shortish = 2 * time.Second
		long     = 10 * time.Minute
		longer   = 120 * time.Hour
	)

	b.Run("Marshalling JSON", func(b *testing.B) {
		log := lg.Extend()

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			log.Info("Test",
				lg.F{"uint", 10},
				lg.F{"uint", []int{0, 2, 7, 12, 22}},
				lg.F{"int", -100},
				lg.F{"int", []int{-10, -4, 2, 7, 12, 22}},
				lg.F{"float", 3.14159},
				lg.F{"floats", []float64{2.5, 7.5, 12.4, 22.3}},
				lg.F{"string", "hello world"},
				lg.F{"strings", []string{"foo", "bar", "baz"}},
				lg.F{"time", now},
				lg.F{"times", []time.Time{before, now, later}},
				lg.F{"duration", longer},
				lg.F{"durations", []time.Duration{short, shortish, long, longer}},
				lg.F{"error", err},
				lg.F{"extended1", "💩"},
				lg.F{"extended2", "🤔"},
				lg.F{"extended3", "🙊"})
		}
	})

	b.Run("Marshalling JSON parallel", func(b *testing.B) {
		log := lg.Extend()

		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				log.Info("Test",
					lg.F{"uint", 10},
					lg.F{"uint", []int{0, 2, 7, 12, 22}},
					lg.F{"int", -100},
					lg.F{"int", []int{-10, -4, 2, 7, 12, 22}},
					lg.F{"float", 3.14159},
					lg.F{"floats", []float64{2.5, 7.5, 12.4, 22.3}},
					lg.F{"string", "hello world"},
					lg.F{"strings", []string{"foo", "bar", "baz"}},
					lg.F{"time", now},
					lg.F{"times", []time.Time{before, now, later}},
					lg.F{"duration", longer},
					lg.F{"durations", []time.Duration{short, shortish, long, longer}},
					lg.F{"error", err},
					lg.F{"extended1", "💩"},
					lg.F{"extended2", "🤔"},
					lg.F{"extended3", "🙊"})
			}
		})
	})
}
