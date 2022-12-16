package common

import "fmt"

type Range struct {
	lower int
	upper int
}

func (r Range) Unpack() []int {
	x := r.upper - r.lower + 1
	a := make([]int, x)
	for i := range a {
		a[i] = r.lower + i
	}
	return a
}

func (r Range) Lower() int {
	return r.lower
}

func (r Range) Upper() int {
	return r.upper
}

func (r1 Range) Contains(r2 Range) bool {
	return r2.lower >= r1.lower && r2.upper <= r1.upper
}

func (r Range) ContainsValue(value int) bool {
	return value >= r.lower && value <= r.upper
}

func (r1 Range) Intersection(r2 Range) *Range {
	var min, max Range
	if r1.upper < r2.lower {
		min = r1
	} else {
		min = r2
	}
	if (r1 == min) {
		max = r2
	} else {
		max = r1
	}
	//min ends before max starts -> no intersection
    if min.upper < max.lower {
		return nil //the ranges don't intersect
	}
	var upper int
    if (min.upper < max.upper)  {
		upper = min.upper
	} else {
		upper = max.upper
	}
    return &Range{max.lower , upper}
}

func (r1 Range) Union(r2 Range) []Range {
	if r1.Contains(r2){
		return []Range{r1}		
	}  else if r2.Contains(r1) {
		return []Range{r2}		
	} else if r1.ContainsValue(r2.lower) && !r1.ContainsValue(r2.upper){
		return []Range{{r1.lower, r2.upper}}		
	} else if r1.ContainsValue(r2.upper) && !r1.ContainsValue(r2.lower){
		return []Range{{r2.lower, r1.upper}}		
	} else if r2.ContainsValue(r1.lower) && !r2.ContainsValue(r1.upper){
		return []Range{{r2.lower, r1.upper}}		
	} else if r2.ContainsValue(r1.upper) && !r2.ContainsValue(r1.lower){
		return []Range{{r1.lower, r2.upper}}		
	} else {
		return []Range{r1, r2}
	}
}

func (r Range) Size() int {
	return r.upper - r.lower + 1
}

func (r Range) String() string {
	return fmt.Sprintf("%v", r.Unpack())
}

func NewRange(lower, upper int) Range {
	if lower > upper {
		return Range{upper, lower}
	} else {
		return Range{lower, upper}
	}
	
}