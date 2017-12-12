package json_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/autopilothq/lg/encoding/json"
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
