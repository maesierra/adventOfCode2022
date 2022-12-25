package day25

import (
	"strconv"

	"maesierra.net/advent-of-code/2022/common"
)

type Day25 struct {
}

var digitsToDecimal = map[string]int {
	"0": 0,
	"1": 1,
	"2": 2,
	"-": -1,
	"=": -2,
}

var digitsToSnafu = map[string]string {
	"0": "0",
	"1": "1",
	"2": "2",
	"3": "1=",
	"4": "1-",
	"5": "10",
	"6": "11",
	"7": "12",
	"8": "2-",
	"9": "2=",
}

var snafuAdd = map[string]map[string]string {
	"0":  {"0": "0",  "1": "1",  "2": "2",  "1=": "1=", "1-": "1-"},
	"1":  {"0": "1",  "1": "2",  "2": "1=", "1=": "1-", "1-": "10"},
	"2":  {"0": "2",  "1": "1=", "2": "1-", "1=": "10", "1-": "11"},
	"1=": {"0": "1=", "1": "1-", "2": "10", "1=": "11", "1-": "12"},
	"1-": {"0": "1-", "1": "10", "2": "11", "1=": "12", "1-": "2="},
	"10": {"0": "10", "1": "11", "2": "12", "1=": "2=", "1-": "2-"},
	"11": {"0": "11", "1": "12", "2": "2=", "1=": "2-", "1-": "30"},
	"12": {"0": "12", "1": "2=", "2": "2-", "1=": "30", "1-": "31"},
	"2=": {"0": "2=", "1": "2-", "2": "30", "1=": "31", "1-": "32"},
}


func (d Day25) DecimalToSnafu(n int) string {
	res := ""
	decimal := n
	carry := ""
	for ; ; {
		digit := digitsToSnafu[strconv.Itoa(decimal % 5)]
		if carry != "" {
			digit = snafuAdd[digit][carry]
			carry = ""
		}
		if len(digit) > 1 {
			carry = digit[:1]
			res = digit[1:] + res
		} else {
			res = digit + res 
		}
		if decimal == 0 {
			break
		}
		decimal /= 5
	}
	if res[0] == '0' {
		return res[1:]
	} else {
		return res
	}
}

func (d Day25) SnafuToDecimal(n string) int {
	multiplier := 1
	decimal := 0
	for i := len(n) - 1; i >= 0; i-- {
		decimal += digitsToDecimal[string(n[i])] * multiplier
		multiplier *= 5
	}
	return decimal
}

func (d Day25) SolvePart1(inputFile string, data []string) string {
	
	var sum int = 0
	for _, line := range common.ReadFileIntoLines(inputFile) {
		sum += d.SnafuToDecimal(line)
	}
	return d.DecimalToSnafu(sum)

}

func (d Day25) SolvePart2(inputFile string, data []string) string {
	var score int = 0
	return strconv.Itoa(int(score))
}
