package day1

import (
	"container/heap"
	"strconv"
	"strings"

	"maesierra.net/advent-of-code/2022/common"
)

type Day1 struct {
}

func (d Day1) SolvePart1(inputFile string) string {
	input := strings.Split(common.ReadFile(inputFile), "\n\n")
	var max int64 = 0
	for _, perElf := range input {
		var total int64 = d.SumCalories(perElf)
		if total > max {
			max = total
		}
	}
	return strconv.Itoa(int(max))

}

func (d Day1) SolvePart2(inputFile string) string {
	input := strings.Split(common.ReadFile(inputFile), "\n\n")
	top := &common.MaxIntHeap{0, 0, 0}
	heap.Init(top)

	for _, perElf := range input {
		heap.Push(top, int(d.SumCalories(perElf)))
	}
	top3 := (*top)[0] + (*top)[1] + (*top)[2]
	return strconv.Itoa(int(top3))
}

func (Day1) SumCalories(perElf string) int64 {
	var total int64 = 0
	for _, i := range strings.Split(perElf, "\n") {
		c, err := strconv.ParseInt(i, 10, 0)
		common.PanicIfError(err)
		total += c
	}
	return total
}
