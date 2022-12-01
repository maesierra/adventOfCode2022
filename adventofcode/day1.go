package adventofcode

import (
	"crypto/md5"
	"fmt"

	"maesierra.net/advent-of-code/2022/common"
)

type Day1 struct {
}

func (d Day1) SolvePart1(inputFile string) string {
	input := common.ReadFile(inputFile)
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))

}

func (d Day1) SolvePart2(inputFile string) string {
	input := common.ReadFile(inputFile) + "..93287529752"
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}
