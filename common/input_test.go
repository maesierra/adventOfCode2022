package common

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_readFile(t *testing.T) {
	contents := "aaabbcc"
	file, error := ioutil.TempFile(os.TempDir(), "sample")
	Check(error)
	defer os.Remove(file.Name())
	_, error = file.WriteString(contents)
	Check(error)

	if got := ReadFile(file.Name()); got != contents {
		t.Errorf("Hello() = %q, want %q", got, contents)
	}
}
