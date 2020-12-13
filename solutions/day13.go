package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"math"
	"strconv"
	"strings"
)

type Day13Solution struct{
	arrival int64
	busIDs []int64
}

func (s *Day13Solution) Prepare(input string) {
	lines := strings.Split(input, "\n")
	ids := strings.Split(lines[1], ",")
	s.busIDs = make([]int64, len(ids))

	s.arrival = util.MustParseInt(lines[0])
	for k, v := range ids {
		if v == "x" {
			s.busIDs[k] = -1
			continue
		}

		s.busIDs[k] = util.MustParseInt(v)
	}
}

func (s *Day13Solution) Part1() string {
	times := make([]int64, len(s.busIDs))
	copy(times, s.busIDs)

	arrivalReady := 0
	for {
		arrivalReady = 0

		for k, v := range times {
			if v == -1 || v >= s.arrival {
				arrivalReady++
				continue
			}

			times[k] += s.busIDs[k]
		}

		if arrivalReady == len(s.busIDs) {
			break
		}
	}

	var closest int64 = math.MaxInt64
	n := -1
	for k,v := range times {
		if v == -1 {
			continue
		}

		if (v - s.arrival) < (closest - s.arrival) {
			closest = v
			n = k
		}
	}

	return strconv.FormatInt(s.busIDs[n] * (closest - s.arrival), 10)
}

func (s *Day13Solution) CompatibleWithAll(root int64, past []int64) bool {
	for k,v := range past {
		if (root + int64(k)) % v != 0 {
			return false
		}
	}

	return true
}

func (s *Day13Solution) RecursiveTickFinding(passed, root, adjustor int64, past, nums []int64) int64 {
	x := util.TernaryInt64(nums[0] == -1, 1, nums[0])
	past = append(past, x)

	var n int64 = 0
	var first int64
	nIncrements, compat := int64(0), 0
	for {
		if s.CompatibleWithAll(root+(n*adjustor), past) {
			compat++

			if compat == 2 {
				break
			} else if len(nums) == 1 {
				return root+(n*adjustor)
			} else {
				first = root+(n*adjustor)
			}
		}

		if compat > 0 {
			nIncrements++
		}

		n++
	}

	//return s.RecursiveTickFinding(cTicks * n, passed + 1, root, nums[1:])
	return s.RecursiveTickFinding(passed + 1, first, nIncrements * adjustor, past, nums[1:])
}

func (s *Day13Solution) Part2() string {
	n := s.RecursiveTickFinding(1, s.busIDs[0], s.busIDs[0], []int64{ s.busIDs[0] }, s.busIDs[1:])
	//var n int64 = 1
	//maxK, ticksSinceLastMax, currentTickRate := 0, int64(0), int64(1)
	//for {
	//	root := n * s.busIDs[0]
	//
	//	OK := true
	//	for k,v := range s.busIDs {
	//		if k == 0 || v == -1 {
	//			continue
	//		}
	//
	//		r := (root + int64(k)) % v
	//		if r != 0 {
	//			if k > maxK {
	//				maxK = 0
	//				currentTickRate = ticksSinceLastMax
	//			}
	//
	//			n += currentTickRate
	//			ticksSinceLastMax += currentTickRate
	//			OK = false
	//			break
	//		}
	//	}
	//
	//	if OK {
	//		break
	//	}
	//}

	// x + n = x * m
	// let m be the bus ID
	// determine n where x_i-1 and x_i+1 = i +/- 1
	//times := make([]int64, len(s.busIDs))
	//
	//var t int64 = 0
	//for {
	//	times[0] += s.busIDs[0]
	//
	//	for k := range times[1:] {
	//		if s.busIDs[k+1] == -1 {
	//			times[k+1] = times[k]
	//			continue
	//		}
	//
	//		for times[k] > times[k+1] {
	//			times[k+1] += s.busIDs[k+1]
	//		}
	//	}
	//
	//	validTime := true
	//	for k := range times[1:] {
	//		if s.busIDs[k+1] == -1 {
	//			continue
	//		}
	//
	//		expectedTime := times[0]+int64(k)+1
	//
	//		if expectedTime != times[k+1] {
	//			validTime = false
	//			break
	//		}
	//	}
	//
	//	if validTime {
	//		t = times[0]
	//		break
	//	}
	//}

	return strconv.FormatInt(n, 10)
}
