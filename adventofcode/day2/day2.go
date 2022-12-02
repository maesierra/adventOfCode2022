package day2

import (
	"fmt"
	"strconv"
	"strings"

	"maesierra.net/advent-of-code/2022/common"
)

type Day2 struct {
}

var translate map[string]string = map[string]string {
	"A" : "Rock",
	"B" : "Paper",
	"C":  "Scissors",
	"X" : "Rock",
	"Y" : "Paper",
	"Z":  "Scissors",
}

var points map[string]map[string]int = map[string]map[string]int{
	"Rock": map[string]int{
		"Rock":     3 + 1, //Rock vs Rock(1)
		"Paper":    6 + 2, //Rock loooses against Paper(2)
		"Scissors": 0 + 3, //Rock beats Scissors(3)
	},
	"Paper": map[string]int{
		"Rock":     0 + 1, //Paper beats Rock(1)
		"Paper":    3 + 2, //Paper vs Paper(2)
		"Scissors": 6 + 3, //Paper loooses against Scissors (3)
	},
	"Scissors": map[string]int{
		"Rock":     6 + 1, //Scissors loooses against Rock(1)
		"Paper":    0 + 2, //Scissors beats Paper(2)
		"Scissors": 3 + 3, //Scissors vs Scissors (3)
	},
}

func (d Day2) SolvePart1(inputFile string) string {
	
	input := strings.Split(common.ReadFile(inputFile), "\n")
	var score int = 0
	for _, game := range input {
		opponentMove := translate[game[0:1]]
		myMove := translate[game[2:3]]
		res := points[opponentMove][myMove]
		fmt.Printf("%s vs %s => %d\n", opponentMove, myMove, res)
		score += res
	}
	return strconv.Itoa(int(score))

}

func (d Day2) SolvePart2(inputFile string) string {
	input := strings.Split(common.ReadFile(inputFile), "\n")
	strategy := map[string]map[string]string{
        "Rock": map[string]string{
            "X": "Scissors", //Scissors loooses against Rock(
            "Y": "Rock",     //Rock vs Rock
			"Z": "Paper",    //Paper beats Rock
        },
        "Paper": map[string]string{
            "X": "Rock",     //Rock loooses against Paper
            "Y": "Paper",    //Paper vs Paper
			"Z": "Scissors", //Scissors beats Paper
        },
        "Scissors": map[string]string{
            "X": "Paper",    //Paper loooses against Scissors(
            "Y": "Scissors", //Scissors vs Scissors
			"Z": "Rock",     //Rock beats Scissors
        },
    }

	var score int = 0
	for _, game := range input {
		opponentMove := translate[game[0:1]]
		myMove := strategy[opponentMove][game[2:3]]
		res := points[opponentMove][myMove]
		fmt.Printf("%s vs %s => %d\n", opponentMove, myMove, res)
		score += res
	}
	return strconv.Itoa(int(score))
}
