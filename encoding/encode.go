package encoding

import (
	"fmt"
	"reflect"
	"time"

	"github.com/autopilothq/lg/encoding/json"
)

// Encode json encodes a map of key/value pairs
func Encode(enc Encoder, kv map[string]interface{}) (err error) {
	for k, v := range kv {
		if err = enc.AddKey(k); err != nil {
			return err
		}

		if err = EncodeValue(enc, v); err != nil {
			return err
		}
	}

	return nil
}

// EncodeKeyValue json encodes a single key/value pair
func EncodeKeyValue(enc Encoder, key string, value interface{}) (err error) {
	err = enc.AddKey(key)
	if err != nil {
		return err
	}

	return EncodeValue(enc, value)
}

func encodeArray(enc Encoder, arr []interface{}) (err error) {
	if err = enc.StartArray(); err != nil {
		return err
	}

	for _, val := range arr {
		if err = EncodeValue(enc, val); err != nil {
			return err
		}
	}

	return enc.EndArray()
}

func encodeObject(enc Encoder, obj map[string]interface{}) (err error) {
	if err = enc.StartObject(); err != nil {
		return err
	}

	for key, val := range obj {
		if err = EncodeKeyValue(enc, key, val); err != nil {
			return err
		}
	}

	return enc.EndObject()
}

// EncodeValue json encodes a single json value
func EncodeValue(enc Encoder, value interface{}) (err error) {
	switch val := value.(type) {

	case uint:
		return enc.AddUint64(uint64(val))

	case *uint:
		return enc.AddUint64(uint64(*val))

	case uint16:
		return enc.AddUint16(val)

	case *uint16:
		return enc.AddUint16(*val)

	case uint32:
		return enc.AddUint32(val)

	case *uint32:
		return enc.AddUint32(*val)

	case uint64:
		return enc.AddUint64(val)

	case *uint64:
		return enc.AddUint64(*val)

	case int:
		return enc.AddInt64(int64(val))

	case *int:
		return enc.AddInt64(int64(*val))

	case int16:
		return enc.AddInt16(val)

	case *int16:
		return enc.AddInt16(*val)

	case int32:
		return enc.AddInt32(val)

	case *int32:
		return enc.AddInt32(*val)

	case int64:
		return enc.AddInt64(val)

	case *int64:
		return enc.AddInt64(*val)

	case float32:
		return enc.AddFloat32(val)

	case *float32:
		return enc.AddFloat32(*val)

	case float64:
		return enc.AddFloat64(val)

	case *float64:
		return enc.AddFloat64(*val)

	case []uint16:
		return enc.AddArrayish(json.Uint16s(val))

	case []uint32:
		return enc.AddArrayish(json.Uint32s(val))

	case []uint64:
		return enc.AddArrayish(json.Uint64s(val))

	case []int16:
		return enc.AddArrayish(json.Int16s(val))

	case []int32:
		return enc.AddArrayish(json.Int32s(val))

	case []int64:
		return enc.AddArrayish(json.Int64s(val))

	case []float32:
		return enc.AddArrayish(json.Float32s(val))

	case []float64:
		return enc.AddArrayish(json.Float64s(val))

	case []bool:
		return enc.AddArrayish(json.Bools(val))

	case bool:
		return enc.AddBool(val)

	case *bool:
		return enc.AddBool(*val)

	case []byte:
		return enc.AddBytes(val)

	case string:
		return enc.AddString(val)

	case []string:
		return enc.AddArrayish(json.Strings(val))

	case error:
		return enc.AddString(val.Error())

	case nil:
		return enc.AddNull()

	case time.Time:
		return enc.AddTime(val)

	case []time.Time:
		return enc.AddArrayish(json.Times(val))

	case time.Duration:
		return enc.AddDuration(val)

	case []time.Duration:
		return enc.AddArrayish(json.Durations(val))

	case []interface{}:
		return encodeArray(enc, val)

	case map[string]interface{}:
		return encodeObject(enc, val)

	default:
		// TODO remove this debug logging
		fmt.Println("FALLBACK", getTypeName(val), val)
		return enc.AddReflected(val)
	}
}

func getTypeName(i interface{}) string {
	t := reflect.TypeOf(i)

	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}

	return t.Name()
}
