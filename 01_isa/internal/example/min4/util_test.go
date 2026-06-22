package main

func min4Reference(first, second, dst *[4]int32) {
	for i := range dst {
		dst[i] = min(first[i], second[i])
	}
}

func filledSlice[T any](element T, size int) []T {
	result := make([]T, 0, size)

	for i := 0; i < size; i++ {
		result = append(result, element)
	}

	return result
}
