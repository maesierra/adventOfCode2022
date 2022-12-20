package day20

import (
	"fmt"
	"strconv"

	"maesierra.net/advent-of-code/2022/common"
)

type Day20 struct {
}

var logLevel = 0

type Node struct {
	value int
	left *Node
	right *Node
}

func (n1 *Node) AddToRight(value int) *Node {
	n2 := Node{value: value, left: n1}
	n1.right = &n2
	return &n2
}

type List struct {
	start *Node
	markedNode *Node
	order []*Node
	size int
}


func (l *List) Move(n *Node) {
	value := n.value
	if value == 0 {
		return
	}
	steps := value % (l.size - 1)
	var dest *Node = nil
	if steps > 0 {
		//to the right
		dest = n.right
		for i := 0; i < steps - 1; i++ {
			dest = dest.right
		}
	} else {
		//to the left
		dest = n.left
		for i := 0; i < common.IntAbs(steps); i++ {
			dest = dest.left
		}
	}
	if logLevel > 1 {
		fmt.Printf("MoveV3 %d => %d at %d\n", n.value, steps, dest.value)
	}
	//Remove the node from its original position value = value % (l.size - 1)
	n.left.right = n.right
	n.right.left = n.left
	
	//Insert at its new position
	tmp := dest.right
	dest.right = n
	n.left = dest
	n.right = tmp
	tmp.left = n
}




func (l List) String() string {
	str := fmt.Sprintf("[%d", l.start.value)
	for current := l.start.right; current != l.start; current = current.right {
		str += fmt.Sprintf(",%d", current.value)
	}
	return str + "]"
}

func (d Day20) ParseInput(inputFile string, multiplier int) List {
	list := List{
		order: []*Node{},
	}
	var last *Node = nil
	for idx, line := range common.ReadFileIntoLines(inputFile) {
		value, _ := strconv.Atoi(line)
		value *= multiplier
		if idx == 0 {
			list.start = &Node{value: value}
			list.order = append(list.order, list.start)
			last = list.start
		} else {
			last = last.AddToRight(value)
			list.order = append(list.order, last)
		}
		if value == 0 {
			list.markedNode = last
		}
	}
	//Close the list
	last.right = list.start
	list.start.left = last
	list.size = len(list.order)
	return list
}

func (d Day20) SolvePart1(inputFile string, data []string) string {
	
	list := d.ParseInput(inputFile, 1)
	for _, node := range list.order {
		list.Move(node)
		if logLevel > 0 {
			fmt.Printf("Move %d\n%v\n", node.value, list)
		} else {
			fmt.Printf("Move %d\n", node.value)
		}
	}
	sum := 0
	current := list.markedNode
	for i := 0; i < 3000; i++ {
		current = current.right
		if i == 999 || i == 1999 {
			sum += current.value
		}
	}
	sum += current.value
	return strconv.Itoa(sum)

}

func (d Day20) SolvePart2(inputFile string, data []string) string {
	list := d.ParseInput(inputFile, 811589153)
	for i := 0; i < 10; i++ {			
		for _, node := range list.order {
			list.Move(node)
			if logLevel > 0 {
				fmt.Printf("Move %d\n%v\n", node.value, list)
			} else {
				fmt.Printf("Move %d\n", node.value)
			}
		}
	}
	sum := 0
	current := list.markedNode
	for i := 0; i < 3000; i++ {
		current = current.right
		if i == 999 || i == 1999 {
			sum += current.value
		}
	}
	sum += current.value
	return strconv.Itoa(sum)
}
