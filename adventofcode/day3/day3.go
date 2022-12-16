package day3

import (
	"strconv"
	"unicode"

	"maesierra.net/advent-of-code/2022/common"
)

type Day3 struct {
}

func (Day3) calculatePriority(shared rune) int {
	if unicode.IsLower(shared) {
		return int(shared) - int(rune('a')) + 1
	} else {
		return int(shared) - int(rune('A')) + 27
	}
}


func (d Day3) SolvePart1(inputFile string, data []string) string {
	input := common.ReadFileIntoLines(inputFile)
	var sum int = 0
	for _, rucksack := range input {
		len := len(rucksack)
		half := len / 2
		compartment1 := rucksack[0:half]
		compartment2 := rucksack[half:len]
		chars := make(map[rune]bool)
		for _, ch := range compartment1 {
			chars[ch] = true
		}
		var shared rune
		for _, ch := range compartment2 {
			if chars[ch] {
				shared = ch
				break
			}
		}
		sum += d.calculatePriority(shared)
	}

	return strconv.Itoa(int(sum))

}


func (d Day3) SolvePart2(inputFile string, data []string) string {
	groups := common.ReadFileIntoChuncks(inputFile, 3)
	var sum int = 0
	for _, group := range groups {
		histogram := make(map[rune]int)
		for _, rucksack := range group {
			chars := make(map[rune]bool)
			for _, ch := range rucksack {
				if chars[ch] {
					continue
				}
				chars[ch] = true				
				count, present := histogram[ch]
				if (present) {
					histogram[ch] = count + 1 
				} else {
					histogram[ch]  = 1
				}				
			}
		}
		for ch, count := range histogram{
			if count == 3 {
				sum += d.calculatePriority(ch)	
				break
			}
		}
	}

	return strconv.Itoa(int(sum))
}

