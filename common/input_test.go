package common

import (
	"testing"
)

func Test_readFile(t *testing.T) {
	contents := "aaabbcc"
	file := CreateTempFile(contents)

	if got := ReadFile(file.Name()); got != contents {
		t.Errorf("Hello() = %q, want %q", got, contents)
	}
}
