package day10

import (
	"fmt"
	"strconv"
	"strings"

	"maesierra.net/advent-of-code/2022/common"
)

type Day10 struct {
}


type Op struct {
	opType string
	value int	
	nCycles int
}

func (op *Op) Tick() {
	if (op.nCycles > 0) {
		op.nCycles--
	}
}

func (op Op) IsCompleted() bool {
	return op.nCycles == 0
}

func (op Op) apply(register int) int {
	if op.opType == "addx" {
		return register + op.value
	} else {
		return register
	}
}

func (op Op) String() string {
	return fmt.Sprintf("%s %d", op.opType, op.value)
}


func (d Day10) SolvePart1(inputFile string) string {
	
	ops := d.Parse(inputFile)
	controlCycles := map[int]bool{		
		20: true, 
		60: true, 
		100: true, 
		140: true, 
		180: true, 
		220: true, 
	}
	nCycles := 220
	signalStrengths := 0
	running := common.Stack[Op]{}
	register := 1
	for cycle := 1; cycle <= nCycles; cycle++ {
		if _, ok := controlCycles[cycle]; ok {			
			signalStrengths += cycle * register
			fmt.Printf("Signal at %d => %v\n", cycle, register)				
		}
		if !running.IsEmpty() {
			//Something is running
			op := running.TopItem()
			op.Tick()
			if op.IsCompleted() {
				register = op.apply(register)
				fmt.Printf("%d Running %v => %v\n", cycle, op, register)				
				running.PopItem()
			}
		} else {
			op, error := ops.PopItem()
			if (error == nil) {
				fmt.Printf("%d Beginning %v => %v\n", cycle, op, register)
				op.Tick()
				if op.IsCompleted() {
					register = op.apply(register)
					fmt.Printf("%d Running %v => %v\n", cycle, op, register)				
				} else {
					running.PutItem(*op)
				}				
			}
		}		


	}
	return strconv.Itoa(signalStrengths)

}

type Screen struct {
	lines []string
	spritePos int
	width int
}

func (s *Screen) Draw(cycle int) {
	nLine := int((cycle - 1) / s.width)
	pos := len(s.lines[nLine])
	char := "."
	if pos >= s.spritePos - 1 && pos <= s.spritePos + 1 {
		char = "#"
	} 
	s.lines[nLine] += char
}

func (s Screen) String() string {
	str := ""
	for i, line := range s.lines {
		if (i != 0) {
			str += "\n"
		}
		str += line
	}
	return str
}


func (d Day10) SolvePart2(inputFile string) string {
	ops := d.Parse(inputFile)
	screen := Screen{[]string{"","","","","","",}, 1, 40}
	running := common.Stack[Op]{}
	
	for cycle := 1; !ops.IsEmpty(); cycle++ {
		screen.Draw(cycle)
		fmt.Printf("Cycle: %v\n%v\n%s\n", cycle, screen, strings.Repeat("=", 40))
		if !running.IsEmpty() {
			//Something is running
			op := running.TopItem()
			op.Tick()
			if op.IsCompleted() {
				screen.spritePos = op.apply(screen.spritePos)
				running.PopItem()
			}
		} else {
			op, error := ops.PopItem()
			if (error == nil) {
				op.Tick()
				if op.IsCompleted() {
					screen.spritePos = op.apply(screen.spritePos)
				} else {
					running.PutItem(*op)
				}				
			}
		}				
	}	
	return screen.String()

}

func (d Day10) Parse(inputFile string) common.Stack[Op] {
	ops := common.Stack[Op]{}
	for _, line := range common.Reverse(common.ReadFileIntoLines(inputFile)) {
		parts := strings.Split(line, " ")
		value := 0
		if (len(parts) == 2) {
			value, _ = strconv.Atoi(parts[1])
		}
		opType := parts[0]
		nCycles := 1
		if opType == "addx" {
			nCycles = 2
		} 
		ops.PutItem(Op{opType, value, nCycles})
	}
	return ops
}