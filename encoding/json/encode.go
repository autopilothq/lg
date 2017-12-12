package json

import (
	"fmt"
	"reflect"
	"time"
)

// Encode json encodes a map of key/value pairs
func Encode(encoder *Encoder, kv map[string]interface{}) (err error) {
	for k, v := range kv {
		if err = encoder.AddKey(k); err != nil {
			return err
		}

		if err = EncodeValue(encoder, v); err != nil {
			return err
		}
	}

	return nil
}

// EncodeKeyValue json encodes a single key/value pair
func EncodeKeyValue(encoder *Encoder, key string, value interface{}) (err error) {
	err = encoder.AddKey(key)
	if err != nil {
		return err
	}

	return EncodeValue(encoder, value)
}

// EncodeValue json encodes a single json value
func EncodeValue(encoder *Encoder, value interface{}) (err error) {
	switch val := value.(type) {

	case uint:
		return encoder.AddUint64(uint64(val))

	case *uint:
		return encoder.AddUint64(uint64(*val))

	case uint16:
		return encoder.AddUint16(val)

	case *uint16:
		return encoder.AddUint16(*val)

	case uint32:
		return encoder.AddUint32(val)

	case *uint32:
		return encoder.AddUint32(*val)

	case uint64:
		return encoder.AddUint64(val)

	case *uint64:
		return encoder.AddUint64(*val)

	case int:
		return encoder.AddInt64(int64(val))

	case *int:
		return encoder.AddInt64(int64(*val))

	case int16:
		return encoder.AddInt16(val)

	case *int16:
		return encoder.AddInt16(*val)

	case int32:
		return encoder.AddInt32(val)

	case *int32:
		return encoder.AddInt32(*val)

	case int64:
		return encoder.AddInt64(val)

	case *int64:
		return encoder.AddInt64(*val)

	case float32:
		return encoder.AddFloat32(val)

	case *float32:
		return encoder.AddFloat32(*val)

	case float64:
		return encoder.AddFloat64(val)

	case *float64:
		return encoder.AddFloat64(*val)

	case []uint16:
		return encoder.AddArrayish(Uint16s(val))

	case []uint32:
		return encoder.AddArrayish(Uint32s(val))

	case []uint64:
		return encoder.AddArrayish(Uint64s(val))

	case []int16:
		return encoder.AddArrayish(Int16s(val))

	case []int32:
		return encoder.AddArrayish(Int32s(val))

	case []int64:
		return encoder.AddArrayish(Int64s(val))

	case []float32:
		return encoder.AddArrayish(Float32s(val))

	case []float64:
		return encoder.AddArrayish(Float64s(val))

	case []bool:
		return encoder.AddArrayish(Bools(val))

	case bool:
		return encoder.AddBool(val)

	case *bool:
		return encoder.AddBool(*val)

	case []byte:
		return encoder.AddBytes(val)

	case string:
		return encoder.AddString(val)

	case []string:
		return encoder.AddArrayish(Strings(val))

	case error:
		return encoder.AddString(val.Error())

	case nil:
		return encoder.AddNull()

	case time.Time:
		return encoder.AddTime(val)

	case []time.Time:
		return encoder.AddArrayish(Times(val))

	case time.Duration:
		return encoder.AddDuration(val)

	case []time.Duration:
		return encoder.AddArrayish(Durations(val))

	case map[string]interface{}:
		return encoder.AddObject(val)

	default:
		fmt.Println("FALLBACK", getTypeName(val), val)
		return encoder.AddReflected(val)
	}
}

func getTypeName(i interface{}) string {
	t := reflect.TypeOf(i)

	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}

	return t.Name()
}
