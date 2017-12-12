package json_test

import (
	"encoding/json"
	"fmt"
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/autopilothq/lg/encoding/json"
)

var _ = Describe("log encoding Encoder", func() {
	var (
		enc *Encoder
	)

	BeforeEach(func() {
		enc = NewEncoder()
	})

	Describe("AddKey()", func() {
		It("adds a key", func() {
			Expect(enc.AddKey("key")).To(Succeed())
			Expect(enc.String()).To(Equal(`"key":`))
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
			Expect(enc.String()).To(Equal(`"NaN"`))
		})

		It("adds a +Infinity value (32bit)", func() {
			val := float32(math.Inf(1))
			Expect(enc.AddFloat32(val)).To(Succeed())
			Expect(enc.String()).To(Equal(`"+Inf"`))
		})

		It("adds a -Infinity value (32bit)", func() {
			val := float32(math.Inf(-1))
			Expect(enc.AddFloat32(val)).To(Succeed())
			Expect(enc.String()).To(Equal(`"-Inf"`))
		})

		It("adds a float64", func() {
			Expect(enc.AddFloat64(3.5)).To(Succeed())
			Expect(enc.String()).To(Equal(`3.5`))
		})

		It("adds a NaN value (64bit)", func() {
			Expect(enc.AddFloat64(math.NaN())).To(Succeed())
			Expect(enc.String()).To(Equal(`"NaN"`))
		})

		It("adds a +Infinity value (64bit)", func() {
			Expect(enc.AddFloat64(math.Inf(1))).To(Succeed())
			Expect(enc.String()).To(Equal(`"+Inf"`))
		})

		It("adds a -Infinity value (64bit)", func() {
			Expect(enc.AddFloat64(math.Inf(-1))).To(Succeed())
			Expect(enc.String()).To(Equal(`"-Inf"`))
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
			Expect(enc.String()).To(Equal(`"wut!?"`))
		})

		It("escapes double quotes", func() {
			Expect(enc.AddString(`"wut!?"`)).To(Succeed())
			Expect(enc.String()).To(Equal(`"\"wut!?\""`))
		})

		It("escapes back slashes", func() {
			Expect(enc.AddString(`\wut!?\`)).To(Succeed())
			Expect(enc.String()).To(Equal(`"\\wut!?\\"`))
		})

		It("escapes \\n", func() {
			Expect(enc.AddString("\nwut!?\n")).To(Succeed())
			Expect(enc.String()).To(Equal(`"\nwut!?\n"`))
		})

		It("escapes \\r", func() {
			Expect(enc.AddString("\rwut!?\r")).To(Succeed())
			Expect(enc.String()).To(Equal(`"\rwut!?\r"`))
		})

		It("escapes \\t", func() {
			Expect(enc.AddString("\twut!?\t")).To(Succeed())
			Expect(enc.String()).To(Equal(`"\twut!?\t"`))
		})
	})

	Describe("AddObject", func() {
		It("Adds a single element object", func() {
			Expect(enc.AddObject(map[string]interface{}{
				"foo": 1234,
			}))
			Expect(enc.String()).To(Equal(`{"foo":1234}`))
		})

		It("Adds a multiple element object", func() {
			obj := map[string]interface{}{
				"foo": 1234,
				"bar": "wut",
				"baz": 3.5,
			}

			Expect(enc.AddObject(obj))
			finalObj := make(map[string]interface{})
			err := json.Unmarshal(enc.Bytes(), &finalObj)
			Expect(err).To(Succeed())

			mapEquals(finalObj, obj)
		})

		It("Adds nested objects", func() {
			obj := map[string]interface{}{
				"foo": map[string]interface{}{
					"nested": "stuff",
					"more":   1234,
				},
				"bar": "wut",
				"baz": 3.5,
			}

			Expect(enc.AddObject(obj))
			finalObj := make(map[string]interface{})
			err := json.Unmarshal(enc.Bytes(), &finalObj)
			Expect(err).To(Succeed())

			mapEquals(finalObj, obj)
		})
	})

	Describe("AddArray", func() {
		It("Adds a single element array", func() {
			Expect(enc.AddArray([]interface{}{1234}))
			Expect(enc.String()).To(Equal(`[1234]`))
		})

		It("Adds a multiple element object", func() {
			arr := []interface{}{
				"foo", 1234, "bar", "wut", "baz", 3.5,
			}

			Expect(enc.AddArray(arr))
			finalArr := make([]interface{}, 0)
			err := json.Unmarshal(enc.Bytes(), &finalArr)
			Expect(err).To(Succeed())
			arrayEquals(finalArr, arr)
		})

		It("Adds nested objects", func() {
			arr := []interface{}{
				map[string]interface{}{
					"nested": "stuff",
					"more":   1234,
				},
				"bar", "wut", "baz", 3.5,
			}

			Expect(enc.AddArray(arr))
			finalArr := make([]interface{}, 0)
			err := json.Unmarshal(enc.Bytes(), &finalArr)
			Expect(err).To(Succeed())
			arrayEquals(finalArr, arr)
		})
	})
})

func arrayEquals(a1 []interface{}, a2 []interface{}) {
	Expect(a1).To(HaveLen(len(a2)))

	for k, v := range a1 {
		switch v := v.(type) {
		case map[string]interface{}:
			v2, ok := a2[k].(map[string]interface{})
			Expect(ok).To(BeTrue(), "a2 was not a JSON object")
			mapEquals(v, v2)

		case []interface{}:
			v2, ok := a2[k].([]interface{})
			Expect(ok).To(BeTrue(), "a2 was not a JSON Array")
			arrayEquals(v, v2)

		case float64:
			vf, ok := a2[k].(float64)
			if ok {
				Expect(v).To(Equal(vf))
				continue
			}

			// Deal with json.Unmarshal decoding ints into float values
			var vi int
			vi, ok = a2[k].(int)
			if ok {
				Expect(v).To(Equal(float64(vi)))
				continue
			}

			Expect(ok).To(BeTrue(),
				fmt.Sprintf("a2 was not a number: %v", a2[k]))

		default:
			Expect(v).To(Equal(a2[k]),
				fmt.Sprintf("Key `%s` did not match", k))
		}
	}
}

func mapEquals(m1 map[string]interface{}, m2 map[string]interface{}) {
	Expect(m1).To(HaveLen(len(m2)))
	for k, v := range m1 {
		switch v := v.(type) {
		case map[string]interface{}:
			v2, ok := m2[k].(map[string]interface{})
			Expect(ok).To(BeTrue(), "m2 was not a JSON object")
			mapEquals(v, v2)

		case []interface{}:
			v2, ok := m2[k].([]interface{})
			Expect(ok).To(BeTrue(), "m2 was not a JSON Array")
			arrayEquals(v, v2)

		case float64:
			vf, ok := m2[k].(float64)
			if ok {
				Expect(v).To(Equal(vf))
				continue
			}

			// Deal with json.Unmarshal decoding ints into float values
			var vi int
			vi, ok = m2[k].(int)
			if ok {
				Expect(v).To(Equal(float64(vi)))
				continue
			}

			Expect(ok).To(BeTrue(),
				fmt.Sprintf("m2 was not a number: %v", m2[k]))

		default:
			Expect(v).To(Equal(m2[k]),
				fmt.Sprintf("Key `%s` did not match", k))
		}
	}
}
