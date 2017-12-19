package benchmark

import (
	"bytes"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/autopilothq/lg"
)

func BenchmarkFieldLogging(b *testing.B) {
	var out bytes.Buffer
	lg.RemoveOutput(os.Stdout)
	lg.AddOutput(&out, lg.JSON())
	defer lg.RemoveOutput(&out)

	var (
		now   = time.Now()
		err   = errors.New("It broke")
		bools = []bool{
			true, false, true, false, true, false, true, false, true, false}
		ints      = []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
		uints     = []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		floats    = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		strings   = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		durations = []time.Duration{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		times     = []time.Time{
			time.Unix(0, 0),
			time.Unix(1, 0),
			time.Unix(2, 0),
			time.Unix(3, 0),
			time.Unix(4, 0),
			time.Unix(5, 0),
			time.Unix(6, 0),
			time.Unix(7, 0),
			time.Unix(8, 0),
			time.Unix(9, 0),
		}
	)

	b.Run("Marshalling JSON", func(b *testing.B) {
		log := lg.Extend()

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			log.Info("Test",
				lg.F{"uint", 10},
				lg.F{"uint", uints},
				lg.F{"int", -100},
				lg.F{"int", ints},
				lg.F{"uint", uint(5)},
				lg.F{"uint", uints},
				lg.F{"float", 3.14159},
				lg.F{"floats", floats},
				lg.F{"string", "hello world"},
				lg.F{"strings", strings},
				lg.F{"bool", true},
				lg.F{"bools", bools},
				lg.F{"time", now},
				lg.F{"times", times},
				lg.F{"duration", time.Duration(0)},
				lg.F{"durations", durations},
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
					lg.F{"uint", uints},
					lg.F{"int", -100},
					lg.F{"int", ints},
					lg.F{"uint", uint(5)},
					lg.F{"uint", uints},
					lg.F{"float", 3.14159},
					lg.F{"floats", floats},
					lg.F{"string", "hello world"},
					lg.F{"strings", strings},
					lg.F{"bool", true},
					lg.F{"bools", bools},
					lg.F{"time", now},
					lg.F{"times", times},
					lg.F{"duration", time.Duration(0)},
					lg.F{"durations", durations},
					lg.F{"error", err},
					lg.F{"extended1", "💩"},
					lg.F{"extended2", "🤔"},
					lg.F{"extended3", "🙊"})
			}
		})
	})
}
