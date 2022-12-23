package day23

import (
	"fmt"
	"math"
	"strconv"

	"gonum.org/v1/gonum/mat"
	"maesierra.net/advent-of-code/2022/common"
	"maesierra.net/advent-of-code/2022/common/intmatrix"
)

var logLevel int = 0

type Day23 struct {
}

type Direction struct {
	direction string
	options []string
}

var directions [4]Direction = [4]Direction {
	{direction: "N", options: []string{"N", "NE", "NW"}},
	{direction: "S", options: []string{"S", "SE", "SW"}},
	{direction: "W", options: []string{"W", "NW", "SW"}},
	{direction: "E", options: []string{"E", "NE", "SE"}},
}

type Map struct {
	m *mat.Dense
	elves map[int]*Elf
	direction int

}

func (m *Map) Expand() {
	row := m.m.RowView(0)
	for c := 0; c < row.Len(); c++ {
		if row.AtVec(c) != 0 {
			m.m = intmatrix.AddRowToTheTop(1, m.m)
			for _, e := range m.elves {
				e.position.Y++
			}
			break
		}
	}
	
	column := m.m.ColView(0)
	for c := 0; c < column.Len(); c++ {
		if column.AtVec(c) != 0 {
			m.m = intmatrix.AddColumnToTheLeft(1, m.m)
			for _, e := range m.elves {
				e.position.X++
			}
			break
		}
	}
	row = m.m.RowView(m.m.RawMatrix().Rows -1)
	for c := 0; c < row.Len(); c++ {
		if row.AtVec(c) != 0 {
			m.m = intmatrix.AddRowToTheBottom(1, m.m)
			break
		}
	}
	
	column = m.m.ColView(m.m.RawMatrix().Cols - 1)
	for c := 0; c < column.Len(); c++ {
		if column.AtVec(c) != 0 {
			m.m = intmatrix.AddColumnToTheRight(1, m.m)
			break
		}
	}
}

func (m Map) IsOccupied(elf Elf, direction string) bool{
	var x int = -1
	var y int = -1
	switch direction {
	case "N":
		y = elf.position.Y - 1
		x = elf.position.X		
	case "NE":
		y = elf.position.Y - 1
		x = elf.position.X + 1
	case "E":
		y = elf.position.Y
		x = elf.position.X + 1
	case "SE":
		y = elf.position.Y + 1
		x = elf.position.X + 1
	case "S":
		y = elf.position.Y + 1
		x = elf.position.X
	case "SW":
		y = elf.position.Y + 1
		x = elf.position.X - 1
	case "W":
		y = elf.position.Y
		x = elf.position.X- 1
	case "NW":
		y = elf.position.Y - 1
		x = elf.position.X - 1
	}
	return m.m.At(y, x) != 0
}

func (m Map) CanMove(elf Elf) bool {
	for y := elf.position.Y - 1; y <= elf.position.Y + 1; y++ {
		for x := elf.position.X - 1; x <= elf.position.X + 1; x++ {
			if x == elf.position.X && y == elf.position.Y {
				continue
			}
			if m.m.At(y, x) != 0 {
				return true
			}
		}
	}
	//No other Elves are in one of those eight positions
	return false
}

func (m Map) CalculateMovements() {
	used := map[string]*Elf{}
	for _, elf := range m.elves {
		if !m.CanMove(*elf) {
			continue
		}
		for idx := range directions {
			d := directions[(m.direction + idx) % len(directions)]
			free := true
			for _, option := range d.options  {
				if m.IsOccupied(*elf, option) {
					free = false
					break
				}
			}
			if free {
				point := elf.ProposeMove(d.direction)
				otherElf, present := used[fmt.Sprintf("%v", point)]
				if present {
					otherElf.movement = nil
				} else {
					elf.movement = &point						
				}
				used[fmt.Sprintf("%v", point)] = elf					
				break
			}
		}		
	}
}

func (m Map) MoveElves() int{
	count := 0
	for _, elf := range m.elves {
		if elf.movement == nil {
			continue
		}
		count++
		m.m.Set(elf.position.Y, elf.position.X, 0)
		m.m.Set(elf.movement.Y, elf.movement.X, float64(elf.id))
		elf.position = * elf.movement
		elf.movement = nil
	}
	return count
}

func (m *Map) Round() int{
	//Check if we need to expand
	m.Expand()
	//Calculate the movements
	m.CalculateMovements()
	//Execute
	movements := m.MoveElves()
	//Change directions order
	m.direction = (m.direction + 1) % len(directions)
	return movements
}

func (m Map) String() string{
	str := ""
	for i := 0; i < m.m.RawMatrix().Rows; i++ {
		if i != 0 {
			str += "\n"
		}
		for j := 0; j < m.m.RawMatrix().Cols; j++ {
			if m.m.At(i, j) == 0 {
				str += "."
			} else {
				str +=  "#"
			}
		}		
	}
	return str
}

type Elf struct {
	id int
	position common.Point
	movement *common.Point
}

func (elf Elf) ProposeMove(direction string) common.Point{
	var x int
	var y int
	switch direction {
	case "N":
		y = elf.position.Y - 1
		x = elf.position.X		
	case "E":
		y = elf.position.Y
		x = elf.position.X + 1
	case "S":
		y = elf.position.Y + 1
		x = elf.position.X
	case "W":
		y = elf.position.Y
		x = elf.position.X- 1
	}
	return common.Point{X: x, Y: y}
}

func (e Elf) String() string {
	return fmt.Sprintf("[%v]%v", e.id, e.position)
}

func (d Day23) ParseInput(inputFile string) Map {
	lines := common.ReadFileIntoLines(inputFile)
	rows := len(lines)
	columns := len(lines[0])
	data := make([]float64, rows * columns)
	for i := range data {
		data[i] = 0
	}
	
	elves := map[int]*Elf{}
	m := mat.NewDense(rows, columns, data)
	for row, line :=  range lines {
		for column, ch := range line {
			if ch == '#' {
				elf := Elf {
					id: len(elves) + 1,
					position: common.Point{X: column, Y: row},
				}
				m.Set(row, column, float64(elf.id))
				elves[elf.id] = &elf
			}
		}
	}
	return Map{
		m: m, 
		elves: elves,
		direction: 0,
	}
}



func (d Day23) SolvePart1(inputFile string, data []string) string {
	
	m  := d.ParseInput(inputFile)
	if logLevel > 1 {
		fmt.Printf("== Initial State ==\n%v\n", m)
	}			 
	for round := 1; round <= 10; round++ {
		m.Round()
		if logLevel > 1 {
			fmt.Printf("== End of Round %d ==\n%v\n", round, m)
		}
	}
	//One final expand to make the calculations easier
	m.Expand()
	if logLevel > 0 {
		fmt.Printf("== Final ==\n%v\n", m)
	}
	count := (m.m.RawMatrix().Rows - 2) * (m.m.RawMatrix().Cols - 2) - len(m.elves)
	return strconv.Itoa(count)

}

func (d Day23) SolvePart2(inputFile string, data []string) string {
	m  := d.ParseInput(inputFile)
	if logLevel > 1 {
		fmt.Printf("== Initial State ==\n%v\n", m)
	}	
	movements := math.MaxInt	
	round := 0	 
	for ; movements != 0 ; round++ {
		movements = m.Round()
		if logLevel > 1 {
			fmt.Printf("== End of Round %d ==\n%v\n", round, m)
		} else {
			fmt.Printf("Round %d movements: %d\n", round, movements)
		}
	}
	return strconv.Itoa(round)
}
