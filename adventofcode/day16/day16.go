package day16

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/albertorestifo/dijkstra"
	"maesierra.net/advent-of-code/2022/common"
)

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

type ValveSim struct {
	root            *Valve
	valves          map[string]*Valve
	connectionGraph dijkstra.Graph
	maxMinutes      int
}

type Path struct {
	nodes     []string
	open      map[string]int
	completed bool
}

func (p Path) Key() string {
	str := ""
	for _, item := range p.nodes {
		str = str + item
		openAt, present := p.open[item]
		if present {
			str += fmt.Sprintf("(%d)", openAt)
		}
		str += " "
	}
	return strings.TrimSpace(str)
}

func (p Path) Last() string {
	return p.nodes[len(p.nodes)-1]
}

func (p Path) IsOpen(node string) bool {
	_, present := p.open[node]
	return present
}

func (p Path) Minutes() int {
	return len(p.nodes) + len(p.open)
}

func (d Day16) Parse(inputFile string) ValveSim {
	valveMap := ValveSim{
		nil,
		map[string]*Valve{},
		dijkstra.Graph{},
		30,
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
		if valveMap.root == nil {
			valveMap.root = valve
		}
	}
	return valveMap
}

func (v ValveSim) CalculateOutput(open map[string]int) int {
	output := 0
	for valve, minute := range open {
		output += v.valves[valve].rate * minute
	}
	return output
}
func (v ValveSim) MaxOutputInPath(path Path, position, minute int) []Path {
	if minute > v.maxMinutes {
		return []Path{{path.nodes, path.open, false}}
	}
	if len(path.nodes)-1 == position {
		if minute == v.maxMinutes {
			//No time to open the valve
			return []Path{{path.nodes, path.open, false}}
		}
		open := common.CopyMap(path.open)
		open[path.nodes[position]] = v.maxMinutes - minute
		return []Path{{path.nodes, open, false}}
	}
	valve := v.valves[path.nodes[position]]
	//Check without opening
	res := v.MaxOutputInPath(Path{path.nodes, path.open, false}, position+1, minute+1)
	if valve.rate > 0 && minute <= v.maxMinutes-2 {
		//Check spending an extra minute opening it
		open := common.CopyMap(path.open)
		open[path.nodes[position]] = v.maxMinutes - minute
		res = append(res, v.MaxOutputInPath(Path{path.nodes, open, false}, position+1, minute+2)...)
	}
	return res
}

func (d Day16) SolvePart1(inputFile string, data []string) string {
	simulation := d.Parse(inputFile)
	var max int = 0
	rootPath := Path{[]string{simulation.root.name}, map[string]int{}, false}
	allPaths := map[string]*Path{rootPath.Key(): &rootPath}
	hasMore := true

	for iter := 0; hasMore; iter++ {
		hasMore = false
		toProcess := []*Path{}
		for _, path := range allPaths {
			if !path.completed {
				toProcess = append(toProcess, path)
			}
		}
		fmt.Printf("Iter %d processing total %d\n", iter, len(toProcess))
		for idx, path := range toProcess {
			//Find nodes to visit
			//All nodes distinct from the last in the path that are not zero and not already open in the path
			toVisit := []Valve{}
			for name, valve := range simulation.valves {
				if valve.rate > 0 && name != path.Last() && !path.IsOpen(name) {
					toVisit = append(toVisit, *valve)
				}
			}
			for _, valve := range toVisit {
				fmt.Printf("Iter %d processing %d/%d From %v to visit %v\n", iter, idx, len(toProcess), path.Key(), valve.name)
				nodes, _, _ := simulation.connectionGraph.Path(path.Last(), valve.name)
				nodes = append(path.nodes, nodes[1:]...)
				paths := simulation.MaxOutputInPath(Path{nodes, path.open, false}, len(path.nodes), path.Minutes()+1)
				for _, p := range paths {
					allPaths[p.Key()] = &p
					fmt.Printf("Adding %v\n", p.Key())
					hasMore = true
				}
			}
			path.completed = true
		}

		keys := make([]string, 0)
		for key, _ := range allPaths {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		fmt.Printf("Iteration %d\n", iter)
		for _, key := range keys {
			if key != "" {
				fmt.Printf("%v => %d", key, simulation.CalculateOutput(allPaths[key].open))
				if allPaths[key].completed {
					fmt.Printf("[x]")
				}
				fmt.Println()
			}
		}

	}
	fmt.Printf("%v\n", simulation)
	return strconv.Itoa(int(max))

}

func (d Day16) SolvePart2(inputFile string, data []string) string {
	//input := strings.Split(common.ReadFile(inputFile), "\n")
	var score int = 0
	return strconv.Itoa(int(score))
}
