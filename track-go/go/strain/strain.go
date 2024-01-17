package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T ~int | ~[]int | string](arrayInput []T, predicate func(T) bool) []T {
	newArr := make([]T, len(arrayInput))
	c := 0
	for _, v := range arrayInput {
		if predicate(v) {
			newArr[c] = v
			c++
		}
	}
	return newArr[:c]
}

func Discard[T ~int | ~[]int | string](arrayInput []T, predicate func(T) bool) []T {
	return Keep[T](arrayInput, func(v T) bool {
		return !predicate(v)
	})
}
