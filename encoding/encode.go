package encoding

import (
	"reflect"
	"time"

	"github.com/autopilothq/lg/encoding/types"
)

// Encode json encodes a map of key/value pairs
func Encode(enc types.Encoder, kv map[string]interface{}) (err error) {
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
func EncodeKeyValue(enc types.Encoder, key string, value interface{}) (err error) {
	err = enc.AddKey(key)
	if err != nil {
		return err
	}

	return EncodeValue(enc, value)
}

// EncodeStringKeyValue json encodes a single key/string value pair
func EncodeStringKeyValue(enc types.Encoder, key string, value string) (err error) {
	err = enc.AddKey(key)
	if err != nil {
		return err
	}

	return enc.AddString(value)
}

// EncodeTimeKeyValue json encodes a single key/time value pair
func EncodeTimeKeyValue(enc types.Encoder, key string, value time.Time) (err error) {
	err = enc.AddKey(key)
	if err != nil {
		return err
	}

	return enc.AddTime(value)
}

// EncodeAddTimestampKeyValue json encodes a single key/timestamp value pair
func EncodeAddTimestampKeyValue(
	enc types.Encoder, key string,
	value time.Time,
) (err error) {
	err = enc.AddKey(key)
	if err != nil {
		return err
	}

	return enc.AddTimestamp(value)
}

func encodeArray(enc types.Encoder, arr []interface{}) (err error) {
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

func encodeObject(enc types.Encoder, obj map[string]interface{}) (err error) {
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
func EncodeValue(enc types.Encoder, value interface{}) (err error) {
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
		return enc.AddArrayish(types.Uint16s(val))

	case []uint32:
		return enc.AddArrayish(types.Uint32s(val))

	case []uint64:
		return enc.AddArrayish(types.Uint64s(val))

	case []int16:
		return enc.AddArrayish(types.Int16s(val))

	case []int32:
		return enc.AddArrayish(types.Int32s(val))

	case []int64:
		return enc.AddArrayish(types.Int64s(val))

	case []float32:
		return enc.AddArrayish(types.Float32s(val))

	case []float64:
		return enc.AddArrayish(types.Float64s(val))

	case []bool:
		return enc.AddArrayish(types.Bools(val))

	case bool:
		return enc.AddBool(val)

	case *bool:
		return enc.AddBool(*val)

	case []byte:
		return enc.AddBytes(val)

	case string:
		return enc.AddString(val)

	case []string:
		return enc.AddArrayish(types.Strings(val))

	case error:
		return enc.AddString(val.Error())

	case nil:
		return enc.AddNull()

	case time.Time:
		return enc.AddTime(val)

	case []time.Time:
		return enc.AddArrayish(types.Times(val))

	case time.Duration:
		return enc.AddDuration(val)

	case []time.Duration:
		return enc.AddArrayish(types.Durations(val))

	case []interface{}:
		return encodeArray(enc, val)

	case map[string]interface{}:
		return encodeObject(enc, val)

	default:
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
