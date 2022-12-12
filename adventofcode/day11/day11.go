package day11

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"maesierra.net/advent-of-code/2022/common"
)

type Day11 struct {
}

type Monkey struct {
	n int
	items []int
	operation func(i int) int
	testValue int
	ifTrue int
	ifFalse int
	inspectedItems int
}

func (m Monkey) test(item int) bool {
	return item % m.testValue == 0
}

func (m * Monkey) AddItem(item int) {
	m.items = append(m.items, item)
}

func (d Day11) Parse(inputFile string) []Monkey {
	input := strings.Split(common.ReadFile(inputFile), "\n\n")
	monkeys := []Monkey{}
	r, _ := regexp.Compile(`\d+`)
	opRegExp, _ := regexp.Compile(`([*+])|(\d+)`)
	for _, block := range input {
		lines := strings.Split(block, "\n")
		m := r.FindAllStringSubmatch(lines[0], -1)
		nMonkey, _ := strconv.Atoi(m[0][0])
		items := []int{}
		for _, val := range r.FindAllStringSubmatch(lines[1], -1) {
			item, _ := strconv.Atoi(val[0])
			items = append(items, item)
		}
		m = opRegExp.FindAllStringSubmatch(lines[2], -1)
		hasValue := false
		var value int
		if len(m) > 1 {
			hasValue = true
			value, _ = strconv.Atoi(m[1][0])
		}  
		
		var op func(i int) int
		if m[0][0] == "+" {
			if (hasValue) {
				op = func(i int) int {
					return i + value
				}	
			} else {
				op = func(i int) int {
					return i + i
				}	
			}
		} else {
			if (hasValue) {
				op = func(i int) int {
					return i * value
				}
			} else {
				op = func(i int) int {
					return i * i
				}
			}
		}
		m = r.FindAllStringSubmatch(lines[3], -1)
		testValue, _ := strconv.Atoi(m[0][0])
		m = r.FindAllStringSubmatch(lines[4], -1)
		ifTrue, _ := strconv.Atoi(m[0][0])
		m = r.FindAllStringSubmatch(lines[5], -1)
		ifFalse, _ := strconv.Atoi(m[0][0])
		monkey := Monkey{
			nMonkey,
			items,
			op,
			testValue,
			ifTrue,
			ifFalse,
			0,
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func (d Day11) run(monkeys []Monkey, nRounds int, worryAdjustment func (v int) int) int {
	for round := 1; round <= nRounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := &monkeys[i]
			for _, item := range monkey.items {
				monkey.inspectedItems++
				item = worryAdjustment(monkey.operation(item)) 
				if monkey.test(item) {
					monkeys[monkey.ifTrue].AddItem(item)
				} else {
					monkeys[monkey.ifFalse].AddItem(item)
				}
			}
			monkey.items = []int{}
		}
		fmt.Printf("Round %d\n", round)
		for idx, monkey := range monkeys {
			fmt.Printf("Monkey %d: %v\n", idx, monkey.items)
		}
	}
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedItems > monkeys[j].inspectedItems
	})
	monkeyBusiness := monkeys[0].inspectedItems * monkeys[1].inspectedItems
	return monkeyBusiness
}

func (d Day11) SolvePart1(inputFile string) string {	
	monkeys := d.Parse(inputFile)
	monkeyBusiness := d.run(monkeys, 20, func(v int) int {return v / 3})
	return strconv.Itoa(monkeyBusiness)
}

func (d Day11) SolvePart2(inputFile string) string {
	monkeys := d.Parse(inputFile)
	magicProduct := 1
	for _, monkey := range monkeys {
		magicProduct *= monkey.testValue
	}
	monkeyBusiness := d.run(monkeys, 10000, func(v int) int {return v % magicProduct})
	return strconv.Itoa(monkeyBusiness)

}
