package common

func Reverse[A any](s []A) []A {
	items := []A{}
	for i := len(s) - 1; i >= 0; i -= 1 {
		items = append(items, s[i])
	}

	return items
}

func Remove[A any](s []A, i int, zero A) []A{
	s[i] = s[len(s)-1] // Copy last element to index i.
	s[len(s)-1] = zero   // Erase last element (write zero value).
	s = s[:len(s)-1]   // Truncate slice.
	return s
}


func CopyMap[K , V comparable](m map[K]V) map[K]V{
    result := make(map[K]V)
    for k, v := range m {
        result[k] = v
    }
    return result
}