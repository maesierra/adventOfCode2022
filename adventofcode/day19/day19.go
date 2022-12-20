package day19

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"maesierra.net/advent-of-code/2022/common"
)

type Day19 struct {
}

var logLevel = 1

type Resources struct {
	ore int
	clay int
	obsidian int
	geode int
}

func (r Resources) Copy() Resources {
	return Resources{
		r.ore,
		r.clay,
		r.obsidian,
		r.geode,
	}
}

func (r *Resources) Use(cost Cost) {
	r.ore -= cost.oreCost
	if cost.additionalResourceCost > 0 {
		if cost.additionalResource == "clay" {
			r.clay -= cost.additionalResourceCost
		} else {
			r.obsidian -= cost.additionalResourceCost
		}
	}
}

func (r Resources) Get(resource string) int{
	switch resource {
	case "ore":
		return r.ore
	case "clay":
		return r.clay
	case "obsidian":
		return r.obsidian
	case "geode":
		return r.geode
	}
	panic("unknown resource")
}

func (r Resources) String() string {
	return fmt.Sprintf("[ore: %v clay: %v obsidian: %v geode: %v]",r.ore, r.clay, r.obsidian, r.geode)
}

type Robots struct {
	ore int
	clay int
	obsidian int
	geode int
}
func (r Robots) String() string {
	return fmt.Sprintf("[ore: %v clay: %v obsidian: %v geode: %v]",r.ore, r.clay, r.obsidian, r.geode)
}
func (r Robots) Get(resource string) int{
	switch resource {
	case "ore":
		return r.ore
	case "clay":
		return r.clay
	case "obsidian":
		return r.obsidian
	case "geode":
		return r.geode
	}
	panic("unknown resource")
}

var maxGeodes map[int]int = map[int]int{}

type Cost struct {
	oreCost int
	additionalResourceCost int
	additionalResource string
}

func (c Cost) CanAffort(resources Resources) bool{
	if resources.ore < c.oreCost {
		return false
	}
	if c.additionalResourceCost > 0 {
		if c.additionalResource == "clay" {
			return resources.clay >= c.additionalResourceCost
		} else {
			return resources.obsidian >= c.additionalResourceCost
		}
	}
	return true
}
type Blueprint struct {
	id int
	robots map[string]Cost
	maxRequirements map[string]int
}



func (d Day19) Collect(current Resources, robots Robots) Resources{
	current.ore += robots.ore
	current.clay += robots.clay
	current.obsidian += robots.obsidian
	current.geode += robots.geode
	return current
}

func (d Day19) AddNewRobot(robots Robots, resource string) Robots {
	switch resource {
	case "ore":
		robots.ore++
	case "clay":
		robots.clay++
	case "obsidian":
		robots.obsidian++
	case "geode":
		robots.geode++	
	}
	return robots
}

// Return all the posible robots that can be created with the curent resources
func (d Day19) CalculateRobotProduction(blueprint Blueprint, resources Resources, robots Robots, minutes int) []string {
	res := []string{""}
	for r1, c1 := range blueprint.robots {
		if !c1.CanAffort(resources) {
			continue
		}
		if r1 != "geode" {
			//Check if we are going to need any more resources of those
			maxRequired := minutes * blueprint.maxRequirements[r1]
			currentProjection := (robots.Get(r1) * minutes) + resources.Get(r1)
			if currentProjection >= maxRequired {
				//We are already producing enough
				if logLevel > 3 {
					fmt.Printf("(%d) We're are producing already %v of %v\n", minutes, blueprint.maxRequirements[r1], r1)
				}
				continue
			}

		}
		res = append(res, r1)
	}
	//Give priority to productions that will produce something
	sort.SliceStable(res, func(i, j int) bool {
		p1 := res[i]
		p2 := res[j]
		//if we're producting a new geode robot it will have hightest priority
		if p1 == "geode" {
			return true
		} else if p2 == "geode" {
			return false
		}
		if p1 == "obsidian" {
			return true
		} else if p2 == "obsidian" {
			return false
		}
		if p1 == "clay" {
			return true
		} else if p2 == "clay" {
			return false
		}
		//if we have clay or obsidian, it could be better to wait rather than build an ore robot
		if resources.clay > 0 || resources.obsidian > 0 {
			return false
		}
		
		if p1 == "ore" {
			return true
		} else if p2 == "ore" {
			return false
		}
		return true
	})
	return res
}

func (d Day19) Run(blueprint Blueprint, resources Resources, robots Robots, nextRobot string, minutes int) (Resources, bool) {
	if minutes == 0 {
		return resources, true
	}
	currentMax, present := maxGeodes[minutes] 
	resources = d.Collect(resources, robots)
	if logLevel > 4 {
		fmt.Printf("%d Factory produced %v\n", 25 - minutes, resources)

	}
	nGeode := resources.geode
	if present && currentMax > nGeode {
		return resources, false
	}
	if nGeode > 0 {
		if !present || nGeode > currentMax {
			maxGeodes[minutes] = nGeode
			if logLevel > 2 {
				fmt.Println("New best option")
			}			
		} 
	}
	if nextRobot != "" {
		robots = d.AddNewRobot(robots, nextRobot)
		resources.Use(blueprint.robots[nextRobot])
	}
	if logLevel > 4 {
		fmt.Printf("%d Created robots %v\n", 25 - minutes, robots)
	}
	max := resources
	options := d.CalculateRobotProduction(
		blueprint,
		resources,
		robots,
		minutes,
	)
	for _, p := range options {
		resouces, ok := d.Run(blueprint, resources, robots, p, minutes - 1)
		if ok {
			if resouces.geode > max.geode {				
				max = resouces
			}	
		}
		
	}
	return max, true
}



func (d Day19) ParseInput(inputFile string) []Blueprint {
	res := []Blueprint{}
	r, _ := regexp.Compile(`\d+`) 
	for _, line := range common.ReadFileIntoLines(inputFile) {
		m := r.FindAllStringSubmatch(line, -1)
		id, _ := strconv.Atoi(m[0][0])
		oreRobotOreCost, _ := strconv.Atoi(m[1][0])
		clayRobotOreCost, _ := strconv.Atoi(m[2][0])
		obsidianRobotOreCost, _ := strconv.Atoi(m[3][0])
		obsidianRobotClayCost, _ := strconv.Atoi(m[4][0])
		geodeRobotOreCost, _ := strconv.Atoi(m[5][0])
		geodeRobotObsidianCost, _ := strconv.Atoi(m[6][0])
		maxOre := common.IntMax(oreRobotOreCost, clayRobotOreCost)
		maxOre = common.IntMax(maxOre, obsidianRobotOreCost)
		maxOre = common.IntMax(maxOre, geodeRobotOreCost)
		res = append(res, Blueprint{id, map[string]Cost{
			"ore": {oreCost:  oreRobotOreCost},
			"clay": {oreCost:  clayRobotOreCost},
			"obsidian": {
				oreCost:  obsidianRobotOreCost, 
				additionalResource: "clay", 
				additionalResourceCost: obsidianRobotClayCost},
			"geode":{
				oreCost:  geodeRobotOreCost, 
				additionalResource: "obsidian", 
				additionalResourceCost: geodeRobotObsidianCost},
			}, map[string]int{
				"ore": maxOre,
				"clay": obsidianRobotClayCost,
				"obsidian": geodeRobotObsidianCost,
			}})
	}
	return res
}

func (d Day19) SolvePart1(inputFile string, data []string) string {
	blueprints := d.ParseInput(inputFile)
	total := 0
	for _, blueprint := range blueprints {
		fmt.Printf("Calculating blueprint %d\n", blueprint.id)
		resources := Resources{0, 0, 0, 0}
		robots := Robots{1, 0, 0, 0}
		resources, _ = d.Run(
			blueprint,
			resources,
			robots,
			"",
			24,
		)
		qualityNumber := blueprint.id * resources.geode
		maxGeodes = map[int]int{} //Clean the cache
		total += qualityNumber
	} 
	return strconv.Itoa(total)

}

func (d Day19) SolvePart2(inputFile string, data []string) string {
	blueprints := d.ParseInput(inputFile)
	nBlueprints := 3
	if len(data) > 0 {
		v, _ := strconv.Atoi(data[0])
		nBlueprints = v
	}
	total := 1
	for i := 0; i < nBlueprints; i++ {
		blueprint := blueprints[i]
		fmt.Printf("Calculating blueprint %d\n", blueprint.id)
		resources := Resources{0, 0, 0, 0}
		robots := Robots{1, 0, 0, 0}
		resources, _ = d.Run(
			blueprint,
			resources,
			robots,
			"",
			32,
		)
		total *= resources.geode
		maxGeodes = map[int]int{} //Clean the cache
	} 
	return strconv.Itoa(total)

}


