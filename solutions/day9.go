package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"math"
	"strconv"
)

type Day9Solution struct {
	preambleLength int64
	dataStream []int64
	part1Solve int64
}

func (s *Day9Solution) Prepare(input string) {
	if s.preambleLength == 0 {
		s.preambleLength = 25
	}

	s.part1Solve = -1
	s.dataStream = util.ParseInts(input, "\n")
}

func (s *Day9Solution) SUM2Valid(workingSet []int64, sum int64) [2]int64 {
	for x,a := range workingSet {
		for y,b := range workingSet {
			if x != y && a + b == sum {
				return [2]int64{int64(x), int64(y)}
			}
		}
	}

	return [2]int64{-1,-1}
}

func (s *Day9Solution) Part1() string {
	//workingSet := make([]int64, len(s.dataStream))
	//copy(workingSet, s.dataStream)

	offset := 25
	for s.SUM2Valid(s.dataStream[offset-25:offset], s.dataStream[offset]) != [2]int64{-1,-1} {
		offset++
	}

	s.part1Solve = s.dataStream[offset]
	return strconv.FormatInt(s.dataStream[offset], 10)
}

func (s *Day9Solution) Part2() string {
	if s.part1Solve == -1 {
		s.Part1()
	}

	var sum, slidingBack, slidingFront int64 = s.dataStream[0] + s.dataStream[1], 0, 1
	for sum != s.part1Solve {
		if sum < s.part1Solve {
			slidingFront++
			sum += s.dataStream[slidingFront]
		} else {
			sum -= s.dataStream[slidingBack]
			slidingBack++
		}

		if slidingFront >= int64(len(s.dataStream)) {
			return "FAILED: No Solution"
		}
	}

	var min, max int64 = math.MaxInt64, math.MinInt64
	for slidingBack <= slidingFront {
		val := s.dataStream[slidingBack]

		if val > max {
			max = val
		}

		if val < min {
			min = val
		}

		slidingBack++
	}

	return strconv.FormatInt(min + max, 10)
}
