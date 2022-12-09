package common

func Reverse[A any](s []A) []A {
	items := []A{}
	for i := len(s) - 1; i >= 0; i -= 1 {
		items = append(items, s[i])
	}

	return items
}