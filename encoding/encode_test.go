package encoding_test

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/autopilothq/lg/encoding"
	fancy "github.com/autopilothq/lg/encoding/json"
)

var _ = Describe("log encoding Encoder", func() {
	var (
		enc *fancy.Encoder
	)

	BeforeEach(func() {
		enc = fancy.NewEncoder()
	})

	Describe("AddObject", func() {
		It("Adds a single element object", func() {
			Expect(EncodeValue(enc, map[string]interface{}{
				"foo": 1234,
			})).To(Succeed())
			Expect(enc.String()).To(Equal(`{"foo":1234}`))
		})

		It("Adds a multiple element object", func() {
			obj := map[string]interface{}{
				"foo": 1234,
				"bar": "wut",
				"baz": 3.5,
			}

			Expect(EncodeValue(enc, obj)).To(Succeed())
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

			Expect(EncodeValue(enc, obj)).To(Succeed())
			finalObj := make(map[string]interface{})
			err := json.Unmarshal(enc.Bytes(), &finalObj)
			Expect(err).To(Succeed())

			mapEquals(finalObj, obj)
		})
	})

	Describe("AddArray", func() {
		It("Adds a single element array", func() {
			Expect(EncodeValue(enc, []interface{}{1234})).To(Succeed())
			Expect(enc.String()).To(Equal(`[1234]`))
		})

		It("Adds a multiple element object", func() {
			arr := []interface{}{
				"foo", 1234, "bar", "wut", "baz", 3.5,
			}

			Expect(EncodeValue(enc, arr)).To(Succeed())
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

			Expect(EncodeValue(enc, arr)).To(Succeed())
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
