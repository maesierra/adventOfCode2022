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

func (r1 Range) Contains(r2 Range) bool {
	return r2.lower >= r1.lower && r2.upper <= r1.upper
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

func (r Range) String() string {
	return fmt.Sprintf("%v", r.Unpack())
}

func NewRange(lower, upper int) Range {
	return Range{lower, upper}
}