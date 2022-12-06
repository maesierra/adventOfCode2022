package common

import "fmt"

type RuneHistogram struct {
	set map[rune]int
}

func (r *RuneHistogram) Add(ch rune) {
	c, present := r.set[ch]
	if present {
		r.set[ch] = c + 1	
	} else {
		r.set[ch] = 1
	}
}

func (r *RuneHistogram) Remove(ch rune) {
	c, present := r.set[ch]
	if present && c > 0 {
		r.set[ch] = c - 1	
	}
}

func (r RuneHistogram) ValuesWithCount(count int) []rune {
	res := make([]rune, 0)
	for key,value := range r.set {
		if value == count {
			res = append(res, key)
		}
	}
	return res
}

func (r RuneHistogram) String() string {
	s := ""
	for key,value := range r.set {
		s = s + fmt.Sprintf("%v: %v\n", string(key), value)
	}
	return s
}

func MakeRuneHistogram(size int) RuneHistogram {
	return RuneHistogram{map[rune]int{}}
}