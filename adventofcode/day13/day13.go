package day13

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"maesierra.net/advent-of-code/2022/common"
)

var debug = true

type Day13 struct {
}

type Packet struct {
	isList bool
	intValue int
	items []Packet
}

func NewInt(intValue int) Packet{
	return Packet{false, intValue, []Packet{}}
}

func NewList() Packet {
	return Packet{true, 0, []Packet{}}
}

func (i Packet) asList() Packet {
	if i.isList {
		return i
	} else {
		return Packet{true, 0, []Packet{i}}
	}
}



func (i Packet) Compare(other Packet) int {
	if debug {
		fmt.Printf("Compare %v vs %v...\n", i, other)
	}
	var compare int
	if i.isList {
		if other.isList {
			//Compare 2 lists
			idx := 0
			for ; idx < i.Count(); idx++ {
				if idx >= other.Count()  {
					//other has less elements -> this is bigger
					compare = 1
					break
				}
				compare = i.items[idx].Compare(other.items[idx])
				if compare != 0 {
					break
				} 
			}
			if idx == i.Count() {
				compare = i.Count() - other.Count()
			}
		} else {
			//Compare list with int
			compare = i.Compare(other.asList())
		}
	} else {
		if other.isList {
			//Compare int with list
			compare = i.asList().Compare(other)
		} else {
			//compare 2 ints
			compare = i.intValue - other.intValue
		}
	}
	if debug {
		if compare < 0 {
			fmt.Printf("%v < %v\n", i, other)
		} else if compare > 0 {
			fmt.Printf("%v > %v\n", i, other)
		} else {
			fmt.Printf("%v = %v\n", i, other)
		}	
	}
	return compare
}

func (i Packet) String() string {
	if i.isList {
		str := "["
		for idx, item := range i.items {
			if idx != 0 {
				str += ","
			}
			str += item.String()
		}
		return str + "]"
	} else {
		return strconv.Itoa(i.intValue)
	}
}

func (i Packet) Count() int {
	if i.isList {
		return len(i.items)
	} else {
		return 1
	}
}

func (l *Packet) Add(i Packet) {
	if l.isList {
		l.items = append(l.items, i)
	}
}


func (d Day13) Parse(line string) Packet {
	stack := common.Stack[* Packet]{}
	current := ""
	for _, ch := range line {
		if ch == '[' || ch == ',' || ch == ']' {
			if current != "" {
				value, _ := strconv.Atoi(current)
				item := NewInt(value)	
				stack.TopItem().Add(item)
				current = ""
			}
		}
		switch ch {
		case '[':
			newList := NewList()
			stack.PutItem(&newList)
		case ']': 			
			item, _ := stack.PopItem()
			if stack.Count() > 0 {
				top := stack.TopItem()
				top.Add(**item)
			} else {
				return **item
			}
		case ',':
			continue
		default:
			current += string(ch)
		}
	}
	panic("Malformed list")
}

func (d Day13) SolvePart1(inputFile string, data []string) string {
	
	input := common.ReadFileIntoLBlocks(inputFile, "\n\n")	
	sum := 0
	for i, block := range input {
		lines := strings.Split(block, "\n")
		item1 := d.Parse(lines[0])
		item2 := d.Parse(lines[1])
		fmt.Printf("%d: Checking %v vs %v\n", i + 1, item1, item2)
		if item1.Compare(item2) < 0 {
			fmt.Printf("%d %v < %v\n", i + 1, item1, item2)
			sum += i + 1
		} else {
			fmt.Printf("%d %v > %v\n", i + 1, item1, item2)
		}
	}
	return strconv.Itoa(sum)

}

func (d Day13) SolvePart2(inputFile string, data []string) string {
	input := common.ReadFileIntoLBlocks(inputFile, "\n\n")	
	divider1 := d.Parse("[[2]]")
	divider2 := d.Parse("[[6]]")
	packets := []Packet{divider1, divider2}
	decoderKey := 1
	for _, block := range input {
		lines := strings.Split(block, "\n")		
		packets = append(packets, d.Parse(lines[0]))
		packets = append(packets, d.Parse(lines[1]))
	}
	sort.SliceStable(packets, func(i, j int) bool {
		return packets[i].Compare(packets[j]) < 0
	})	
	for idx, packet := range packets {
		fmt.Printf("%d: %v", idx + 1, packet)
		if packet.Compare(divider1) == 0 || packet.Compare(divider2) == 0 {
			decoderKey *= idx + 1
		}
	}
	return strconv.Itoa(decoderKey)
}
