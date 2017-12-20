package buffer_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/autopilothq/lg/encoding/buffer"
)

var _ = Describe("log encoding Buffer", func() {
	var (
		buf *Buffer
	)

	BeforeEach(func() {
		buf = NewBuffer()
	})

	Describe("AppendByte()", func() {
		It("appends a byte", func() {
			buf.AppendByte('!')
			Expect(buf.Bytes()).To(Equal([]byte{'!'}))
		})

		It("appends multiple bytes", func() {
			buf.AppendByte('w')
			buf.AppendByte('u')
			buf.AppendByte('t')
			buf.AppendByte('!')
			Expect(buf.Bytes()).To(Equal([]byte{'w', 'u', 't', '!'}))
		})
	})

	Describe("Write()", func() {
		var amount int

		BeforeEach(func() {
			var err error
			amount, err = buf.Write([]byte{'w', 'u', 't', '!'})
			Expect(err).To(Succeed())
		})

		It("appends the bytes", func() {
			Expect(buf.Bytes()).To(Equal([]byte{'w', 'u', 't', '!'}))
		})

		It("indicates how many bytes where appended", func() {
			Expect(amount).To(Equal(4))
		})
	})

	Describe("AppendString()", func() {
		It("appends a string", func() {
			buf.AppendString("Wut!")
			Expect(buf.String()).To(Equal("Wut!"))
		})
	})

	Describe("AppendDuration()", func() {
		It("appends a duration", func() {
			buf.AppendDuration(100 * time.Nanosecond)
			Expect(buf.String()).To(Equal("100"))
		})
	})

	Describe("AppendTime()", func() {
		It("appends a time", func() {
			t := time.Date(2017, time.December, 20, 11, 20, 1, 152000000, time.UTC)
			buf.AppendTime(t)
			Expect(buf.String()).To(Equal("2017-12-20T11:20:01.152"))
		})
	})

	Describe("AppendTimestamp()", func() {
		It("appends a time", func() {
			t := time.Date(2017, time.December, 20, 11, 20, 1, 52000, time.UTC)
			buf.AppendTimestamp(t)
			Expect(buf.String()).To(Equal(
				fmt.Sprintf("%d", t.UnixNano())))

		})
	})

	Describe("Append ints", func() {
		It("appends an int64", func() {
			buf.AppendInt(-1234)
			Expect(buf.String()).To(Equal("-1234"))
		})

		It("appends an int16", func() {
			buf.AppendInt16(-1234)
			Expect(buf.String()).To(Equal("-1234"))
		})

		It("appends an int32", func() {
			buf.AppendInt32(-1234)
			Expect(buf.String()).To(Equal("-1234"))
		})
	})

	Describe("AppendPaddedInt", func() {
		It("panics on minWidth == 0", func() {
			Expect(func() {
				buf.AppendPaddedInt(123, 0)
			}).To(Panic())
		})

		It("appends with minWidth == 1", func() {
			buf.AppendPaddedInt(10, 1)
			Expect(buf.String()).To(Equal("10"))
		})

		It("zero-padds with minWidth == 2", func() {
			buf.AppendPaddedInt(5, 2)
			Expect(buf.String()).To(Equal("05"))
		})

		It("zero-padds with minWidth == 3", func() {
			buf.AppendPaddedInt(50, 3)
			Expect(buf.String()).To(Equal("050"))
		})

		It("zero-padds with minWidth == 4", func() {
			buf.AppendPaddedInt(50, 4)
			Expect(buf.String()).To(Equal("0050"))
		})
	})

	Describe("Append uints", func() {
		It("appends an uint64", func() {
			buf.AppendUint(uint64(1234))
			Expect(buf.String()).To(Equal("1234"))
		})

		It("appends an uint16", func() {
			buf.AppendUint16(1234)
			Expect(buf.String()).To(Equal("1234"))
		})

		It("appends an uint32", func() {
			buf.AppendUint32(1234)
			Expect(buf.String()).To(Equal("1234"))
		})
	})

	Describe("Append floats", func() {
		It("appends an float64", func() {
			buf.AppendFloat(3.14159, 64)
			Expect(buf.String()).To(Equal("3.14159"))
		})

		It("appends an float32", func() {
			buf.AppendFloat(-3.14159, 32)
			Expect(buf.String()).To(Equal("-3.14159"))
		})
	})

	Describe("Append bools", func() {
		It("appends true", func() {
			buf.AppendBool(true)
			Expect(buf.String()).To(Equal("true"))
		})

		It("appends fals", func() {
			buf.AppendBool(false)
			Expect(buf.String()).To(Equal("false"))
		})
	})

	Describe("Reset()", func() {
		It("resets the buffer to empty", func() {
			buf.AppendString("Wut!")
			Expect(buf.String()).To(Equal("Wut!"))
			buf.Reset()
			Expect(buf.Len()).To(Equal(0))
			Expect(buf.Bytes()).To(HaveLen(0))
		})
	})
})
