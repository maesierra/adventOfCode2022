package common

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

func ReadFileIntoLines(path string) []string {
	return strings.Split(ReadFile(path), "\n")
}

func ReadFileIntoLBlocks(path, separator string) []string {
	return strings.Split(ReadFile(path), separator)
}

func ReadFileIntoChuncks(path string, chunckSize int) [][]string {
	return ChunkSlice(ReadFileIntoLines(path), chunckSize)
}

func ChunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func ParseInt(s string) int {
	intVal, err := strconv.Atoi(s)
	PanicIfError(err)
	return intVal
}