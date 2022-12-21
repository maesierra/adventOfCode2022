package day21

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"maesierra.net/advent-of-code/2022/common"
)

type Day21 struct {
}

var logLevel = 0

type Monkey struct {
	name string
	value int
	resolved bool
	monkeyType string
	monkey1 *Monkey
	monkey2 *Monkey
	operation string
}

func (m Monkey) String() string {
	str := m.name
	if m.resolved {
		str += "[R]"
	} else if m.CanResolve() {
		str += "[C]"
	}
	if m.monkeyType == "value" || m.resolved {
		str += fmt.Sprintf(": %v", m.value)
	} else {
		str += fmt.Sprintf(": %v %v %v", m.monkey1.name, m.operation, m.monkey2.name)
	}
	return str
}

func (m Monkey) CanResolve() bool {
	if m.monkeyType == "value" {
		return true
	} else {
		return m.monkey1.resolved && m.monkey2.resolved
	}
}

func (m * Monkey) Resolve() {
	m.resolved = true
	if m.monkeyType == "value" {
		return
	}
	switch m.operation {
	case "+":
		m.value = m.monkey1.value + m.monkey2.value
	case "-":
		m.value = m.monkey1.value - m.monkey2.value
	case "*":
		m.value = m.monkey1.value * m.monkey2.value
	case "/":
		m.value = m.monkey1.value / m.monkey2.value		
	}
}

func (d Day21) GetOrCreateMonkey(name string, monkeys map[string]*Monkey) *Monkey {
	monkey, present := monkeys[name]
	if !present {
		monkey = &Monkey{
			name: name,
			resolved: false,
			operation: "",
			monkey1: nil,
			monkey2: nil,
		}
		monkeys[name] = monkey
	}
	return monkey
}

func (d Day21) ParseInput(inputFile string) MonkeyMap {
	r, _ := regexp.Compile(`(.*): (.*)`)
	opRegExp, _ := regexp.Compile(`(.*) ([-+/*]) (.*)`)
	monkeys := MonkeyMap{}
	for _, line := range common.ReadFileIntoLines(inputFile) {
		m := r.FindAllStringSubmatch(line, -1)
		monkey := d.GetOrCreateMonkey(m[0][1], monkeys)		
		part2 := m[0][2]
		m = opRegExp.FindAllStringSubmatch(part2, -1)
		if len(m) > 0 {
			monkey.monkeyType = "operation"
			monkey.operation = m[0][2]
			monkey.monkey1 = d.GetOrCreateMonkey(m[0][1], monkeys)
			monkey.monkey2 = d.GetOrCreateMonkey(m[0][3], monkeys)

		} else {
			monkey.monkeyType = "value"
			monkey.value, _ = strconv.Atoi(part2)
			monkey.resolved = true
		}
	}
	return monkeys
}

type MonkeyMap map[string]*Monkey 

func (m MonkeyMap) String() string {
	str := ""
	for _, monkey := range m {
		str += monkey.String() + "\n"
	}
	return str
}


func (d Day21) SolvePart1(inputFile string, data []string) string {
	
	monkeys := d.ParseInput(inputFile)
	if logLevel > 1 {
		fmt.Printf("%v\n\n\n", monkeys)
	}

	for ;!monkeys["root"].resolved; {
		//Find the next unresolved monkey that can be resoled
		var monkey * Monkey = nil
		for _, m := range monkeys {
			if !m.resolved && m.CanResolve() {
				monkey = m
				break
			}
		}
		if monkey == nil {
			panic("No more candidates. Something is wrong")
		}
		if logLevel > 0 {
			fmt.Printf("Resolving monkey %v\n", monkey.name)
		}
		monkey.Resolve()
		if logLevel > 1 {
			fmt.Printf("%v\n\n\n", monkeys)
		}
			
	}
	return strconv.Itoa(monkeys["root"].value)

}

type Term interface {
	Resolved() bool
	Unknown() bool
	Value() int
	Part1() *Term
	Part2() *Term
	Operation() string
	Name() string
}

type Value struct {
	value int
	name string
}

func (v *Value) Name() string {
	return v.String()
}

func (v *Value) Operation() string {
	return ""
}

func (v *Value) String() string {
	return strconv.Itoa(v.value)
}

func (v *Value) Resolved() bool {
	return true
}

func (v *Value) Unknown() bool {
	return false
}

func (v *Value) Value() int {
	return v.value
}

func (v *Value) Part1() *Term {
	return nil
}
func (v *Value) Part2() *Term {
	return nil
}

type Operation struct {
	name string
	part1 * Term
	part2 * Term
	operation string
}

func (o *Operation) Name() string {
	return o.name
}

func (o *Operation) String() string {
	return fmt.Sprintf("(%v %v %v)", *o.part1, o.operation, *o.part2)
}


func (o *Operation) Operation() string {
	return o.operation
}

func (o *Operation) Resolved() bool {
	return (*o.part1).Resolved() && (*o.part2).Resolved()
}

func (o *Operation) Unknown() bool {
	return false
}

func (o *Operation) Value() int {
	switch o.operation {
	case "+":
		return (*o.part1).Value() + (*o.part2).Value()
	case "-":
		return (*o.part1).Value() - (*o.part2).Value()
	case "*":
		return (*o.part1).Value() * (*o.part2).Value()
	case "/":
		return (*o.part1).Value() / (*o.part2).Value()		
	}
	panic("unknown operation")
}

func (o *Operation) Part1() * Term {
	return o.part1
}
func (o *Operation) Part2() * Term {
	return o.part2
}

type Unknown struct {

}

func (u *Unknown) Name() string {
	return "humn"
}

func (u *Unknown) Operation() string {
	return ""
}


func (u *Unknown) Unknown() bool {
	return true
}

func (u *Unknown) String() string {
	return "humn"
}

func (u *Unknown) Resolved() bool {
	return true
}

func (u *Unknown) Value() int {
	return math.MinInt
}

func (u *Unknown) Part1() *Term {
	return nil
}
func (u *Unknown) Part2() *Term {
	return nil
}


func (m Monkey) ToTerm() *Term {
	if m.monkeyType == "value" {
		if m.name == "humn" {
			unknown := Term(&Unknown{})
			return &unknown
		} else {
			value := Term(&Value{name: m.name, value: m.value})
			return &value
		}		
	}
	op := Term(&Operation{
		name: m.name,
		part1: m.monkey1.ToTerm(),
		part2: m.monkey2.ToTerm(),
		operation: m.operation,
	})
	return &op
}

func (d Day21) ContainsUknown(t Term) bool{
	if t.Unknown() {
		return true
	}
	if t.Part1() == nil {
		return false
	}
	return d.ContainsUknown(*t.Part1()) || d.ContainsUknown(*t.Part2())
}

func (d Day21) InverseTerm(t Term, op string, right Term, inverse bool) Term {
	switch op {
	case "+":
		op := Term(&Operation{
			name: t.Name(),
			part1: &right,
			part2: &t,
			operation: "-",
		})
		return op
	case "*":
		op := Term(&Operation{
			name: t.Name(),
			part1: &right,
			part2: &t,
			operation: "/",
		})
		return op
	case "-":
		if !inverse {
			op := Term(&Operation{
				name: t.Name(),
				part1: &right,
				part2: &t,
				operation: "+",
			})
			return op	
		} else {
			op := Term(&Operation{
				name: t.Name(),
				part1: &t,
				part2: &right,
				operation: "-",
			})
			return op	
		}
	case "/":
		if !inverse {
			op := Term(&Operation{
				name: t.Name(),
				part1: &right,
				part2: &t,
				operation: "*",
			})
			return op	
		} else {
			op := Term(&Operation{
				name: t.Name(),
				part1: &t,
				part2: &right,
				operation: "/",
			})
			return op	
		}
	}
	panic("unknown operation")
}

func (d Day21) SolvePart2(inputFile string, data []string) string {
	monkeys := d.ParseInput(inputFile)
	if logLevel > 1 {
		fmt.Printf("%v\n\n\n", monkeys)
	}
	root := monkeys["root"]
	left := *root.monkey1.ToTerm()
	right := *root.monkey2.ToTerm()
	if !d.ContainsUknown(left) {
		left = right
		right = *root.monkey1.ToTerm()
	}
	if logLevel > 0 {
		fmt.Printf("Iteration start: %v = %v\n", left, right)
	}
	//until left side is the unknown 
	for ;!left.Unknown(); {
		//try to move from left to right side
		if d.ContainsUknown(*left.Part1()) {
			right = d.InverseTerm(*left.Part2(), left.Operation(), right, false)
			left = *left.Part1()
		} else {
			right = d.InverseTerm(*left.Part1(), left.Operation(), right, true)
			left = *left.Part2()
		}
		if logLevel > 0 {
			fmt.Printf("After inverse: %v = %v\n", left, right)
		}
	}
	value := right.Value()
	return strconv.Itoa(value)

}
