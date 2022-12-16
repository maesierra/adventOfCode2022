package day16

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/albertorestifo/dijkstra"
	"maesierra.net/advent-of-code/2022/common"
)

type Day16 struct {
}

type Valve struct {
	name string
	rate int
	connections []*Valve
}

type ValveMap struct {
	root *Valve
	valves map[string]*Valve
	connectionGraph dijkstra.Graph
}
func (d Day16) Parse(inputFile string) ValveMap {
	valveMap := ValveMap{
		nil,
		map[string]*Valve{}, 
		dijkstra.Graph{},
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


func (d Day16) SolvePart1(inputFile string, data []string) string {	
	valveMap := d.Parse(inputFile)
	var max int = 0
	current := valveMap.root.name
	open := map[string]int{}
	minute := 1
	for ;; {
		candidates := map[string]int{}
		for name, valve := range valveMap.valves {
			if name != current {
				path, _, _ := valveMap.connectionGraph.Path(current, valve.name)


				candidates[valve.name] = cost
			}
		}
	}
	fmt.Printf("%v\n", valveMap)
	return strconv.Itoa(int(max))

}

func (d Day16) SolvePart2(inputFile string, data []string) string {
	//input := strings.Split(common.ReadFile(inputFile), "\n")
	var score int = 0
	return strconv.Itoa(int(score))
}
