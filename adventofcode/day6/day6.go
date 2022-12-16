package day6

import (
	"fmt"
	"strconv"
	"strings"

	"maesierra.net/advent-of-code/2022/common"
)

type Day6 struct {
}

func (Day6) detectPattern(inputFile string, size int) string {
	input := strings.TrimSpace(common.ReadFile(inputFile))
	set := common.MakeRuneHistogram(0)
	idx := 0
	for pos, ch := range input {
		if pos < size {
			set.Add(ch)
		} else {
			set.Remove(rune(input[idx]))
			set.Add(ch)
			idx++
		}
		fmt.Printf("%v\n", set)
		if len(set.ValuesWithCount(1)) == size {
			return strconv.Itoa(pos + 1)
		}

	}
	return strconv.Itoa(0)
}



func (d Day6) SolvePart1(inputFile string, data []string) string {	
	return d.detectPattern(inputFile, 4)

}

func (d Day6) SolvePart2(inputFile string, data []string) string {
	return d.detectPattern(inputFile, 14)
}
