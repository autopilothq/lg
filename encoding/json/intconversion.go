package json

func uintArray32To16(in []uint32) *[]uint16 {
	out := make([]uint16, len(in))

	for i, v := range in {
		out[i] = uint16(v)
	}

	return &out
}

func uintArray16To32(in []uint16) []uint32 {
	out := make([]uint32, len(in))

	for i, v := range in {
		out[i] = uint32(v)
	}

	return out
}

func intArray32To16(in []int32) *[]int16 {
	out := make([]int16, len(in))

	for i, v := range in {
		out[i] = int16(v)
	}

	return &out
}

func intArray16To32(in []int16) []int32 {
	out := make([]int32, len(in))

	for i, v := range in {
		out[i] = int32(v)
	}

	return out
}
