package common

import (
	"os"
)

func ReadFile(path string) string {
	data, error := os.ReadFile(path)
	Check(error)
	return string(data)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
