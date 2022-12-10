package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"maesierra.net/advent-of-code/2022/common"
)

type Day5 struct {
	
}


type CrateMover struct {
	crates map[string]common.Stack[string]
	labels [] string
	blockMove bool
}

func (c *CrateMover) AddStack(label string) {
	c.crates[label] = common.Stack[string]{}
	c.labels = append(c.labels, label)
}

func (c CrateMover) Put(label string, item string) {
	crate := c.crates[label]	
	crate.PutItem(item)
	c.crates[label] = crate
}

func (c CrateMover) Move(fromLabel string, toLabel string, count int) {	
	to := c.crates[toLabel]
	from := c.crates[fromLabel]
	items := from.Pop(count, !c.blockMove)
	for _, c := range items {
		to.PutItem(c)
	}
	c.crates[fromLabel] = from
	c.crates[toLabel] = to				
}

func (c CrateMover) String() string {
	str := ""
	for _, label := range c.labels  {
		str += fmt.Sprintf("%s: %v (%d)\n", label, c.crates[label], c.crates[label].Count())
	}
	return str
}

func (c CrateMover) Count() int {
	count := 0
	for _, label := range c.labels  {
		count += c.crates[label].Count()
	}
	return count
}

func (c CrateMover) Get(label string) common.Stack[string] {
	return c.crates[label]
}

func (c CrateMover) Labels() string {
	var result string = ""
	for _, label := range c.labels  {
		result += c.crates[label].TopItem()		
	}
	return result
}

type Movement struct {
	from string
	to string
	count int
}

func (d Day5) parse(inputFile string, version string) (CrateMover, []Movement) {
	input := strings.Split(common.ReadFile(inputFile), "\n")
	crateMover := CrateMover{
		make(map[string]common.Stack[string]),
		make([]string, 0),
		version == "9001",
	}	
	movements := make([]Movement, 0)
	cratesRegExp, _ := regexp.Compile(`\[(.)\]|(    )`)
	numbersRegExp, _ := regexp.Compile(`\d+`) 
	movesRegExp, _ := regexp.Compile(`move (\d+) from (\d+) to (\d+)`) 
	for n, line := range input {
		if movesRegExp.MatchString(line) {
			m := movesRegExp.FindAllStringSubmatch(line, -1)[0]
			count, _ := strconv.Atoi(m[1])
			movements = append(movements, Movement{m[2], m[3], count})
			
		} else if numbersRegExp.MatchString(line) {
			m := numbersRegExp.FindAllStringSubmatch(line, -1)
			for _, i := range m	{
				crateMover.AddStack(i[0])
			}	
			for i := n - 1; i >= 0; i-- {
				m := cratesRegExp.FindAllStringSubmatch(input[i], -1)
				for pos, i := range m	{
					if i[0][0] == '[' {
						crateMover.Put(strconv.Itoa(pos + 1), i[1])
					}
				}
			}			
		} 
	}
	return crateMover, movements
}

func (d Day5) run(crateMover CrateMover, movements []Movement) string{
	fmt.Printf("%v", crateMover)
	for _, movement := range movements {
		fmt.Printf("%v\n", crateMover)
		prevFrom := fmt.Sprintf("%v", crateMover.Get(movement.from))
		prevTo := fmt.Sprintf("%v", crateMover.Get(movement.to))
		moved := common.IntMin(crateMover.Get(movement.from).Count(), movement.count)
		crateMover.Move(movement.from, movement.to, movement.count)							
		fmt.Printf("%d(%d) from %s %v to %s %v\n", movement.count, moved, movement.from, prevFrom, movement.to, prevTo)
		fmt.Printf("%v", crateMover)			 
	}
	return crateMover.Labels()
}

func (d Day5) SolvePart1(inputFile string) string {
	crateMover, movements := d.parse(inputFile, "9000")
	return d.run(crateMover, movements)

}

func (d Day5) SolvePart2(inputFile string) string {
	crateMover, movements := d.parse(inputFile, "9001")
	return d.run(crateMover, movements)
}
