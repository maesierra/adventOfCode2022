package day25_test

import (
	"strconv"
	"testing"

	"maesierra.net/advent-of-code/2022/adventofcode"
	"maesierra.net/advent-of-code/2022/adventofcode/day25"
)

var solver adventofcode.Solver = day25.Day25{}


func ConvertTest(snafu, decimal string, d day25.Day25, t *testing.T) {
	n, _ := strconv.Atoi(decimal)
	got := d.DecimalToSnafu(n)
	if got != snafu {
		t.Errorf("DecimalToSnafu() = %q, want %q", got, snafu)
	}
	got = strconv.Itoa(d.SnafuToDecimal(snafu))
	if got != decimal {
		t.Errorf("SnafuToDecimal() = %q, want %q", got, decimal)
	}
}

func Test_Numbers(t *testing.T) {
	d := day25.Day25{}
	ConvertTest("1=-0-2",    "1747", d, t)
    ConvertTest("12111",     "906", d, t)
    ConvertTest("2=0=",      "198", d, t)
    ConvertTest("21",        "11", d, t)
    ConvertTest("2=01",      "201", d, t)
    ConvertTest("111",       "31", d, t)
    ConvertTest("20012",     "1257", d, t)
    ConvertTest("112",       "32", d, t)
    ConvertTest("1=-1=",     "353", d, t)
    ConvertTest("1-12",      "107", d, t)
    ConvertTest("12",        "7", d, t)
    ConvertTest("1=",        "3", d, t)
    ConvertTest("122",       "37", d, t)
}

func Test_SolvePart1(t *testing.T) {
	input := `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122`
	expected := "2=-1=0"
	adventofcode.SolvePart1Test(input, expected, solver, t)
}

func Test_SolvePart2(t *testing.T) {
	input := `A Y
B X
C Z`
	expected := "12"
	adventofcode.SolvePart2Test(input, expected, solver, t)
}
