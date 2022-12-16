package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"maesierra.net/advent-of-code/2022/common"
)

type Day4 struct {
}

func (d Day4) SolvePart1(inputFile string, data []string) string {
	
	input := strings.Split(common.ReadFile(inputFile), "\n")
	var total int = 0
	for _, assignment := range input {
		elf1, elf2 := d.convertLine(assignment)
		elf1Contained := elf1.Contains(elf2)
		elf2Contained := elf2.Contains(elf1)
		if (elf1Contained || elf2Contained) {
			total++
		}
		fmt.Printf("%v %v %v %v\n", elf1, elf1Contained, elf2, elf2Contained)
	}
	return strconv.Itoa(int(total))

}

func (d Day4) SolvePart2(inputFile string, data []string) string {
	input := strings.Split(common.ReadFile(inputFile), "\n")
	var total int = 0
	for _, assignment := range input {
		elf1, elf2 := d.convertLine(assignment)
		intersects := elf1.Intersection(elf2) != nil
		if (intersects) {
			total++
		}
		fmt.Printf("%v %v %v\n", elf1, elf2, intersects)
	}
	return strconv.Itoa(int(total))
}


func (d Day4) convertLine(s string) (common.Range, common.Range) {
	r, _ := regexp.Compile(`(\d+)-(\d+),(\d+)-(\d+)`)
	m := r.FindAllStringSubmatch(s, -1)[0]
	return common.NewRange(common.ParseInt(m[1]), common.ParseInt(m[2])), common.NewRange(common.ParseInt(m[3]), common.ParseInt(m[4]))
}