package day24

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"

	"maesierra.net/advent-of-code/2022/common"
)

type Day24 struct {
}

var logLevel int = 1

type Map struct {
	rows int
	columns int
	origin common.Point
	destination common.Point
	blizzards []Blizzard	
}

func (m Map) BlizzardsAt(minute int) []Blizzard {
	sizeX := m.columns
	sizeY := m.rows
	res := make([]Blizzard, len(m.blizzards))
	for i, b:= range m.blizzards {
		x := b.pos.X
		y := b.pos.Y
		switch b.direction {
		case ">":
			x = (x + minute) % sizeX
		case "<":
			x = (x - minute) % sizeX
		case "v":
			y = (y + minute) % sizeY
		case "^":
			y = (y - minute) % sizeY	
		}
		if x < 0 {
			x = m.columns + x
		} 
		if y < 0 {
			y = m.rows + y
		}
		res[i].direction = b.direction
		res[i].pos.X = x
		res[i].pos.Y = y
	}
	return res
} 


type Blizzard struct{
	pos common.Point
	direction string
}

type State struct {
	minutes int
	m *Map
	expedition common.Point 
	direction string
	index int
	reverse bool
}

func (s State) Next(point common.Point, direction string) State {
	return State{
		expedition: point,
		m : s.m,
		minutes: s.minutes + 1,
		direction: direction,
		reverse: s.reverse,
	}
}

func (s State) Key() string {
	return fmt.Sprintf("%v-%v", s.expedition, s.minutes)
}
 
func (s State) String() string {
	str := ""
	m := s.m
	blizzards := s.m.BlizzardsAt(s.minutes)
	for row := -1; row <= m.rows; row++ {
		if row != -1 {
			str += "\n"
		}
		for column := -1; column <= m.columns; column++ {			
			ch := "."
			if s.expedition.X == column && s.expedition.Y == row {
				ch = "E"
			} else if m.destination.X == column && m.destination.Y == row {
				ch = "."
			} else if m.origin.X == column && m.origin.Y == row {
				ch = "."
			} else if row == -1 || row == m.rows || column == -1 || column == m.columns {
				ch = "#"
			} else {
				nBlizzards := 0
				for _, b := range blizzards {
					if b.pos.X == column && b.pos.Y == row {
						ch = b.direction
						nBlizzards++
					}
				}	
				if nBlizzards > 1 {
					ch = strconv.Itoa(nBlizzards)
				}
			}			
			str += ch
		}
	}
	return str
}

func (s State) Completed() bool {
	return s.expedition.X == s.m.destination.X && s.expedition.Y == s.m.destination.Y
}

func (s State) Movements() map[string]common.Point {
	res := map[string]common.Point{}
	position := s.expedition
	blizzards := s.m.BlizzardsAt(s.minutes + 1)
	m := s.m
	options := [5]common.Point{
		{X: position.X + 1, Y: position.Y}, 
		{X: position.X, Y: position.Y + 1},
		{X: position.X, Y: position.Y}, //wait option
		{X: position.X - 1, Y: position.Y}, 
		{X: position.X, Y: position.Y - 1}, 
	}
	for idx, p := range options {
		if p.Y < 0 || p.Y >= m.rows || p.X < 0 || p.X >= m.columns {
			//origin and destination are out of the map so we need to check for them
			if (p.X != m.origin.X || p.Y != m.origin.Y) && (p.X != m.destination.X || p.Y != m.destination.Y) {
				continue //Out of map
			}			
		}
		hasBlizzard := false
		for _, b:= range blizzards {
			if b.pos.X == p.X && b.pos.Y == p.Y {
				hasBlizzard = true
				break
			}

		}
		if hasBlizzard {
			continue
		}
		direction := ""
		switch idx {
		case 0:
			direction = ">"
		case 1: 
			direction = "v"
		case 2:
			direction = "-"	
		case 3:
			direction = "<"	
		case 4:
			direction = "^"	
		}
		res[direction] = p
	}
	return res
	
}

type StateQueue []*State

func (sq StateQueue) Len() int { return len(sq) }

func (sq StateQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if sq[i].minutes != sq[j].minutes {
		return sq[i].minutes > sq[j].minutes
	}
	if sq[i].reverse {
		if sq[i].direction == "<" || sq[i].direction == "^" {
			return true
		}
	} else {
		if sq[i].direction == ">" || sq[i].direction == "v" {
			return true
		}
	
	}
	return false
}

func (sq StateQueue) Swap(i, j int) {
	sq[i], sq[j] = sq[j], sq[i]
	sq[i].index = i
	sq[j].index = j
}

func (sq *StateQueue) Push(x any) {
	n := len(*sq)
	item := x.(*State)
	item.index = n
	*sq = append(*sq, item)
}

func (sq *StateQueue) Pop() any {
	old := *sq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*sq = old[0 : n-1]
	return item
}

func (d Day24) ParseInput(inputFile string, reverse bool) State {
	lines := common.ReadFileIntoLines(inputFile)
	m := Map{
		rows: len(lines) - 2,
		columns: len(lines[0]) - 2,
		blizzards: []Blizzard{},
	}
	for row, line := range lines {
		for column, ch := range line {
			x:= column - 1
			y:= row - 1 
			switch ch {
			case '>', '<', '^', 'v':
				m.blizzards = append(m.blizzards, Blizzard{pos: common.Point{X: x, Y: y}, direction: string(ch)})
			case '.': 
				if row == 0 {
					if reverse {
						m.destination = common.Point{X: x, Y: y}
					} else {
						m.origin = common.Point{X: x, Y: y}
					}
				} else if row == len(lines) - 1 {
					if reverse {
						m.origin = common.Point{X: x, Y: y}
					} else {
						m.destination = common.Point{X: x, Y: y}
					}
				}
			}
		}
	}
	return State{
		minutes: 0,
		expedition: m.origin,
		m: &m,
		direction: "-",
		reverse: reverse,
	}
}


func (d Day24) CalculateMinPath(initialState State) int {
	min := math.MaxInt
	candidates := make(StateQueue, 0)
	heap.Init(&candidates)
	heap.Push(&candidates, &initialState)
	visited := map[string]bool{}

	for candidates.Len() > 0 {
		state := heap.Pop(&candidates).(*State)
		visited[state.Key()] = true
		if logLevel > 1 {
			fmt.Printf("Minute %v %v\n%v\n", state.minutes, state.direction, state)
		} else {
			fmt.Printf("Minute %v pos %v %v [%v]\n", state.minutes, state.expedition, state.direction, candidates.Len())
		}
		//end state checks
		if state.Completed() {
			min = common.IntMin(min, state.minutes)
			if logLevel > 0 {
				fmt.Printf("Found the exit in %v\n", state.minutes)
			}
			//Purge
			toRemove := []int{}
			for _, s := range candidates {
				if s.minutes >= state.minutes {
					toRemove = append(toRemove, s.index)
				}
			}
			for _, i := range toRemove {
				heap.Remove(&candidates, i)
			}
			continue
		}
		if min > 0 && state.minutes >= min {
			//not going to improve
			continue
		}

		options := state.Movements()
		if len(options) == 0 {
			if logLevel > 0 {
				fmt.Printf("No options for %v %v \n", state.minutes, state.expedition)
			}
			continue
		}
		for direction, p := range options {
			nextState := state.Next(p, direction)
			_, present := visited[nextState.Key()]
			if !present {
				heap.Push(&candidates, &nextState)
			}
		}

	}
	return min
}

func (d Day24) SolvePart1(inputFile string, data []string) string {

	initialState := d.ParseInput(inputFile, false)	
	
	return strconv.Itoa(d.CalculateMinPath(initialState))

}


func (d Day24) SolvePart2(inputFile string, data []string) string {
	initialStateToExit := d.ParseInput(inputFile, false)
	initialStateFromExit := d.ParseInput(inputFile, true)
	toExit := d.CalculateMinPath(initialStateToExit)
	initialStateFromExit.minutes = toExit
	fromExit := d.CalculateMinPath(initialStateFromExit)
	initialStateToExit.minutes = fromExit
	andBack := d.CalculateMinPath(initialStateToExit)
	return strconv.Itoa(andBack)
}
