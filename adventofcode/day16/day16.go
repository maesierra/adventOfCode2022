package day16

import (
	"container/heap"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/albertorestifo/dijkstra"
	"maesierra.net/advent-of-code/2022/common"
)

var logLevel = 1

type Day16 struct {
}

type Valve struct {
	name        string
	rate        int
	connections []*Valve
}

func (v Valve) String() string {
	return v.name
}

type ValveSystem struct {
	root            *Valve
	valves          map[string]*Valve
	connectionGraph dijkstra.Graph
	contributingValves []*Valve
}

type State struct {
	valveSystem 	     *ValveSystem
	key          string
	open      	  map[string]int
	minutesLeft  int
	minutesLeft2  int
	output       int
	index        int
	last		 string
	last2        string
}

func (s State) Valves() []string {
	valves := []string{}
	for valve, _ := range s.open {
		valves = append(valves, valve)
	}
	sort.SliceStable(valves, func(i, j int) bool {
		return s.open[valves[i]] > s.open[valves[j]]
	})
	return valves
}
func (s State) String() string {
	str := ""
	for _, valve := range s.Valves() {
		str += fmt.Sprintf("%v(%d) ", valve, s.open[valve])
		
	}				
	return fmt.Sprintf("%v => %v at %v(%v) / %v(%v)", strings.TrimSpace(str), s.output, s.last, s.minutesLeft, s.last2, s.minutesLeft2)
}


func (s State) IsOpen(node string) bool {
	_, present := s.open[node]
	return present
}

type StateQueue []*State

func (sq StateQueue) Len() int { return len(sq) }

func (sq StateQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	s1 := sq[i]
	s2 := sq[j]
	return s1.MaxPossibleOutput(common.IntMin(s1.minutesLeft, s1.minutesLeft2)) > s2.MaxPossibleOutput(common.IntMin(s2.minutesLeft, s2.minutesLeft2))
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


func (d Day16) Parse(inputFile string) ValveSystem {
	valveMap := ValveSystem{
		root: nil,
		valves: map[string]*Valve{},
		contributingValves: []*Valve{},
		connectionGraph: dijkstra.Graph{},
	}
	r, _ := regexp.Compile(`^Valve (..).*rate=(\d+);.*(?:valves (.*)|valve (.*))$`)
	for _, line := range common.ReadFileIntoLines(inputFile) {
		m := r.FindAllStringSubmatch(line, -1)
		name := m[0][1]
		var valve *Valve
		var present bool
		valve, present = valveMap.valves[name]
		if !present {
			valve = &Valve{name: name}
			valveMap.valves[name] = valve
		}
		valve.rate, _ = strconv.Atoi(m[0][2])
		if valve.rate > 0 {
			valveMap.contributingValves = append(valveMap.contributingValves, valve)
		}
		valve.connections = []*Valve{}
		x := strings.Split(m[0][3]+m[0][4], ", ")
		connections := map[string]int{}
		for _, other := range x {
			connections[other] = 1
			v, present := valveMap.valves[other]
			if !present {
				newValve := Valve{name: other}
				valveMap.valves[other] = &newValve
				valve.connections = append(valve.connections, &newValve)
			} else {
				valve.connections = append(valve.connections, v)
			}
		}
		valveMap.connectionGraph[name] = connections
		if name == "AA" {
			valveMap.root = valve
		}
	}
	return valveMap
}

func (s State) CalculateOutput() int {
	output := 0
	for valve, minute := range s.open {
		output += s.valveSystem.valves[valve].rate * minute
	}
	return output
}

func (s State) MaxPossibleOutput(minute int) int {
	notOpen := []*Valve{}
	for _, v:= range s.valveSystem.contributingValves {
		if _, present := s.open[v.name] ; !present {
			notOpen = append(notOpen, v)
		}
	}
	sort.SliceStable(notOpen, func(i, j int) bool {
		return notOpen[i].rate > notOpen[j].rate
	})
	m := minute
	output := 0
	for _, v:= range notOpen {
		m -= 2 //let's assume it takes a 2 minutes to move and open. Not real but enough to prune 
		if m <= 0 {
			break
		}
		output += v.rate * m
	}
	return s.CalculateOutput() + output
}

func (v ValveSystem) NewStateAfterOpening(s State, valve string, n, distance int) State {
	openAt := s.minutesLeft - 1 - distance
	if n == 2 {
		openAt = s.minutesLeft2 - 1 - distance
	}
	open := common.CopyMap(s.open)
	open[valve] = openAt
	newState := State{
		valveSystem: s.valveSystem,
		open:   open,
		minutesLeft: openAt,
		key: fmt.Sprintf("%v %v(%v)", s.key, valve, openAt),		
	}
	newState.output = newState.CalculateOutput()
	if n == 1 {
		newState.minutesLeft = openAt
		newState.minutesLeft2 = s.minutesLeft2
		newState.last = valve
		newState.last2 = s.last2
	} else if n == 2 {
		newState.minutesLeft2 = openAt
		newState.minutesLeft = s.minutesLeft
		newState.last2 = valve
		newState.last = s.last
	}
	return newState
}


func (s ValveSystem) CalculateValvesToOpen(valves []*Valve) []map[string]bool {
	if len(valves) == 1 {
		return []map[string]bool{{valves[0].name: true}}
	}
	res := []map[string]bool{}
	for _, r := range s.CalculateValvesToOpen(valves[1:]) {
		m := common.CopyMap(r)
		m[valves[0].name] = true
		res = append(res, m)
		m = common.CopyMap(r)
		m[valves[0].name] = false
		res = append(res, m)
	}
	return res
}

func (s ValveSystem) ContributingValves(valves []string) []*Valve {
	res := []*Valve{}
	for _, name := range valves {
		if s.valves[name].rate > 0 {
			res = append(res, s.valves[name])
		}
	}
	return res
}


func (s ValveSystem) CalculateOptions(state State, open2Valves bool) []State {
	if state.minutesLeft < 2 && state.minutesLeft2 < 2 {
		//It takes at least 2 minutes to move and open
		return []State{}
	}
	//Find nodes to visit
	//All nodes distinct from the last in the path that are not zero and not already open in the path
	available := []Valve{}
	currentKey := state.key
	for _, valve := range s.contributingValves {
		if valve.name != state.last && !state.IsOpen(valve.name) {
			available = append(available, *valve)
		}
	}
	toVisit := [][2]*Valve{}
	for idx1, v1 := range available {
		if open2Valves {
			for idx2, v2 := range available {
				if v1.name == v2.name {
					continue
				}
				toVisit = append(toVisit, [2]*Valve{&available[idx1], &available[idx2]})
			}
		} else {
			toVisit = append(toVisit, [2]*Valve{&available[idx1], nil})
		}
	}
	res := []State{}
	for _, valves := range toVisit {
		valve1 := valves[0]
		valve2 := valves[1]
		if logLevel > 2 && valve2 == nil {
			fmt.Printf("From %v(%v) to visit %v\n", currentKey, state.last, valve1.name)
		} 
		path1, _, _ := s.connectionGraph.Path(state.last, valve1.name)
		path1 = path1[1:]
		toOpenInPath1 := s.CalculateValvesToOpen(s.ContributingValves(path1))
		path2 := []string{}
		toOpenInPath2 := []map[string]bool{}
		if open2Valves {
			path2, _, _ = s.connectionGraph.Path(state.last2, valve2.name)
			path2 = path2[1:]
			toOpenInPath2 = s.CalculateValvesToOpen(s.ContributingValves(path2))
		}
		//merge toOpen
		toOpen := toOpenInPath1
		if open2Valves {
			toOpen = []map[string]bool{}
			toOpenKeys := map[string]bool{}
			for _, m1 := range toOpenInPath1 {
				keys := []string{}
				m := map[string]bool{}
				for key, value := range m1 {
					m[key] = value
					if value {
						keys = append(keys, key)
					}
				}
				for _, m2 := range toOpenInPath2 {
					for key, value := range m2 {
						m[key] = value
						if value {
							keys = append(keys, key)
						}
					}
				}
				key := ""
				sort.Strings(keys)
				last := ""
				for _, k := range keys {
					if k != last {
						key += k
					}
					last = k
				}
				if _, present := toOpenKeys[key] ; !present {
					toOpen = append(toOpen, m)
					toOpenKeys[key] = true
				} 
			}
			
		}
		longestPath := path1
		if len(path2) > len(path1) {
			longestPath = path2
		}
		for _, open := range toOpen {
			distance1 := 0
			distance2 := 0
			newState := state
			for idx := range longestPath {
				v1 := "x"
				v2 := "x"
				if idx < len(path1) {
					v1 = path1[idx]
				}
				if open2Valves && idx < len(path2) {
					v2 = path2[idx]
				}
				openValve1 := false
				openValve2 := false
				if openValve, present := open[v1] ; present && openValve {
					openValve1 = true
				}
				if openValve, present := open[v2] ; present && openValve {
					openValve2 = true
				}
				if openValve1 && openValve2 {
					newState = s.NewStateAfterOpening(newState, v1, 1, distance1 + 1)
					if v1 != v2 {
						newState = s.NewStateAfterOpening(newState, v2, 2, distance2 + 1)
						distance2 = 0
					} else {
						distance2++
					}
					distance1 = 0
				} else if openValve1 && !openValve2 {
					newState = s.NewStateAfterOpening(newState, v1, 1, distance1 + 1)
					distance1 = 0
					distance2++
				} else if !openValve1 && openValve2 {
					newState = s.NewStateAfterOpening(newState, v2, 2, distance2 + 1)
					distance1++ 
					distance2 = 0
				} else {
					distance1++
					distance2++
				}
			}
			if logLevel > 2 && open2Valves {
				fmt.Printf("From %v. %v to visit %v [%v]| %v to visit %v [%v] opening %v.\n", currentKey, state.last, valve1.name, path1, state.last, valve2.name, path2, open)	
			}
			res = append(res, newState)
		}
	}
	return res
}


func (valveSystem ValveSystem) Run(maxMinutes int, open2Valves bool) int {
	var max int = 0
	var maxState *State = nil
	initial := State{
		valveSystem: &valveSystem,
		last: valveSystem.root.name,
		last2: valveSystem.root.name,
		open: map[string]int{}, 
		key: "AA",
		output: 0,
		minutesLeft: maxMinutes,
		minutesLeft2: maxMinutes,
	}
	completed := map[string]bool{}
	candidates := make(StateQueue, 0)
	heap.Init(&candidates)
	heap.Push(&candidates, &initial)

	for candidates.Len() > 0 {
		state := heap.Pop(&candidates).(*State)
		if logLevel > 0 {
			fmt.Printf("Processing %v (total %d)\n", state, len(candidates))
		} else {
			fmt.Printf("Processing...(total %d): Max %v\n", len(candidates), max)
		}
		maxPossibleOutput := state.MaxPossibleOutput(state.minutesLeft)
		if maxPossibleOutput < max {
			if logLevel > 2 {
				fmt.Printf("Discarding because %v is less than %v\n", maxPossibleOutput, max)
			}
			continue
		}
		newStates := valveSystem.CalculateOptions(*state, open2Valves)
		for i := range newStates {
			s := newStates[i]
			if logLevel > 1 {
				fmt.Printf("NewState: %v\n", s)
			}
			if s.output > max  {
				if logLevel > 0 {
					fmt.Printf("New max %v => %v\n", s.key, s.output)
				}
				max = s.output
				maxState = &s			
			}
			if _, present := completed[s.key] ; present {
				if logLevel > 1 {
					fmt.Printf("Already present %v\n", s.key)
				}
			} else {
				if s.minutesLeft < 2 || s.minutesLeft2 < 2 {
					if logLevel > 1 {
						fmt.Printf("Discarding because %v run out of time\n", s.key)
					}
					continue
				}
				//Only bother to add it if still has something to be open or if it could improve the current maximum
				//coult improve the current max
				maxPossibleOutput := state.MaxPossibleOutput(s.minutesLeft)
				if maxPossibleOutput < max {
					if logLevel > 1 {
						fmt.Printf("Discarding because %v is less than %v\n", maxPossibleOutput, max)
					}	
				}
				if len(s.open) < len(valveSystem.contributingValves) || maxPossibleOutput  > max {			
					if logLevel > 2 {
						fmt.Printf("Adding %v\n", s)
					}			
					heap.Push(&candidates, &newStates[i])
					
				}
			}					
		}
		completed[state.key] = true
		if logLevel > 1 {
			fmt.Printf("Completing %v\n", state.key)
		}
	}
	fmt.Printf("Max state %v\n", maxState.key)
	return max
}

func (d Day16) SolvePart1(inputFile string, data []string) string {
	valves := d.Parse(inputFile)
	return strconv.Itoa(valves.Run(30, false))

}

func (d Day16) SolvePart2(inputFile string, data []string) string {
	valves := d.Parse(inputFile)
	return strconv.Itoa(valves.Run(26, true))
}
