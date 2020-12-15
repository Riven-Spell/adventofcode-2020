package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"strconv"
)

type Day15Solution struct {
	starters []int64
}

func (s *Day15Solution) Prepare(input string) {
	s.starters = util.ParseInts(input, ",")
}

func (s *Day15Solution) PlayGame(steps int64) int64 {
	seenNumbers := map[int64]int64{}
	cNum := s.starters[len(s.starters)-1]

	for turn,v := range s.starters[:len(s.starters)-1] {
		seenNumbers[v] = int64(turn) + 1
	}

	turn := int64(len(s.starters)) + 1
	tmpLast := seenNumbers[cNum]
	for turn <= steps {
		//lastSeen := seenNumbers[cNum]
		if tmpLast != 0 {
			cNum = (turn - 1) - tmpLast
		} else {
			if turn == int64(len(s.starters)) + 1 {
				seenNumbers[cNum] = turn - 1
			}

			cNum = 0
		}
		tmpLast = seenNumbers[cNum]
		seenNumbers[cNum] = turn

		turn++
	}

	return cNum
}

func (s *Day15Solution) Part1() string {
	return strconv.FormatInt(s.PlayGame(2020), 10)
}

func (s *Day15Solution) Part2() string {
	return strconv.FormatInt(s.PlayGame(30000000), 10)
}
