package common

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) string {
	data, error := os.ReadFile(path)
	PanicIfError(error)
	return string(data)
}

func CreateTempFile(contents string) *os.File {
	file, error := ioutil.TempFile(os.TempDir(), "sample")
	PanicIfError(error)
	//defer os.Remove(file.Name())
	_, error = file.WriteString(contents)
	PanicIfError(error)
	return file
}

func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}
