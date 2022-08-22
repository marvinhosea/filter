package packages

type AllowedTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | string
}

func Filter[T AllowedTypes](slice1, slice2 []T) []T {
	combinedSlice := append(slice1, slice2...)
	temp := make(map[T]bool)
	var filtered []T
	for _, v := range combinedSlice {
		_, ok := temp[v]
		if !ok {
			temp[v] = true
			filtered = append(filtered, v)
		}
	}
	return filtered
}
