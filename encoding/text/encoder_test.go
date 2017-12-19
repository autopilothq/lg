package text_test

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/autopilothq/lg/encoding/text"
)

var _ = Describe("log encoding Text", func() {
	var (
		enc *Encoder
	)

	BeforeEach(func() {
		enc = NewEncoder()
	})

	Describe("AddKey()", func() {
		It("adds a key", func() {
			Expect(enc.AddKey("key")).To(Succeed())
			Expect(enc.String()).To(Equal(`key:`))
		})
	})

	Describe("Add ints", func() {
		It("adds a int16", func() {
			Expect(enc.AddInt16(12345)).To(Succeed())
			Expect(enc.String()).To(Equal(`12345`))
		})

		It("adds a int32", func() {
			Expect(enc.AddInt32(-12345)).To(Succeed())
			Expect(enc.String()).To(Equal(`-12345`))
		})

		It("adds a int64", func() {
			Expect(enc.AddInt64(-12345)).To(Succeed())
			Expect(enc.String()).To(Equal(`-12345`))
		})

		It("adds a uint16", func() {
			Expect(enc.AddUint16(12345)).To(Succeed())
			Expect(enc.String()).To(Equal(`12345`))
		})

		It("adds a uint32", func() {
			Expect(enc.AddUint32(12345)).To(Succeed())
			Expect(enc.String()).To(Equal(`12345`))
		})

		It("adds a uint64", func() {
			Expect(enc.AddUint64(12345)).To(Succeed())
			Expect(enc.String()).To(Equal(`12345`))
		})
	})

	Describe("Add floats", func() {
		It("adds a float32", func() {
			Expect(enc.AddFloat32(3.5)).To(Succeed())
			Expect(enc.String()).To(Equal(`3.5`))
		})

		It("adds a NaN value (32bit)", func() {
			Expect(enc.AddFloat32(float32(math.NaN()))).To(Succeed())
			Expect(enc.String()).To(Equal(`NaN`))
		})

		It("adds a +Infinity value (32bit)", func() {
			val := float32(math.Inf(1))
			Expect(enc.AddFloat32(val)).To(Succeed())
			Expect(enc.String()).To(Equal(`+Inf`))
		})

		It("adds a -Infinity value (32bit)", func() {
			val := float32(math.Inf(-1))
			Expect(enc.AddFloat32(val)).To(Succeed())
			Expect(enc.String()).To(Equal(`-Inf`))
		})

		It("adds a float64", func() {
			Expect(enc.AddFloat64(3.5)).To(Succeed())
			Expect(enc.String()).To(Equal(`3.5`))
		})

		It("adds a NaN value (64bit)", func() {
			Expect(enc.AddFloat64(math.NaN())).To(Succeed())
			Expect(enc.String()).To(Equal(`NaN`))
		})

		It("adds a +Infinity value (64bit)", func() {
			Expect(enc.AddFloat64(math.Inf(1))).To(Succeed())
			Expect(enc.String()).To(Equal(`+Inf`))
		})

		It("adds a -Infinity value (64bit)", func() {
			Expect(enc.AddFloat64(math.Inf(-1))).To(Succeed())
			Expect(enc.String()).To(Equal(`-Inf`))
		})
	})

	Describe("AddBool()", func() {
		It("adds a true value", func() {
			Expect(enc.AddBool(true)).To(Succeed())
			Expect(enc.String()).To(Equal(`true`))
		})

		It("adds a false value", func() {
			Expect(enc.AddBool(false)).To(Succeed())
			Expect(enc.String()).To(Equal(`false`))
		})
	})

	Describe("AddString()", func() {
		It("adds a string", func() {
			Expect(enc.AddString("wut!?")).To(Succeed())
			Expect(enc.String()).To(Equal(`wut!?`))
		})
	})
})
