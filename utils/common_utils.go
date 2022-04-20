package utils

func ToPointerList[T any](origin []T) []*T {
	var result = make([]*T, len(origin))
	for i := range origin {
		result[i] = &origin[i]
	}

	return result
}

func ToStructList[T any](origin []*T) []T {
	var result = make([]T, len(origin))
	for i := range origin {
		result[i] = *origin[i]
	}

	return result
}
