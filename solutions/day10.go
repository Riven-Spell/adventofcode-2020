package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2020/util"
	"sort"
	"strconv"
)

type adapterRange []int64

func (a adapterRange) Len() int {
	return len(a)
}

func (a adapterRange) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a adapterRange) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type Day10Solution struct {
	adapters adapterRange
}

func (s *Day10Solution) Prepare(input string) {
	s.adapters = util.ParseInts(input)
	sort.Sort(s.adapters)
}

func (s *Day10Solution) Part1() string {
	// diff to count
	diffs := make(map[int]int64)

	last := int64(0)
	for k,v := range s.adapters {
		diffs[int(v - last)]++

		fmt.Println(k, v, last, v - last)

		//for _, a2 := range s.adapters[k+1:] {
		//	if a2 - last > 3 {
		//		break
		//	}
		//
		//	fmt.Println("SUBADAPTATION", a2, last, a2 - last)
		//
		//	diffs[int(a2 - last)]++
		//}

		last = v
	}

	fmt.Println(diffs)

	return strconv.FormatInt(diffs[1] * (diffs[3] + 1), 10)
}

func (s *Day10Solution) RecursivePart2(v int, memo map[int]int64) int64 {
	joltage := s.adapters[v]

	//0 = 1 is already set, so there is no need to memoize it manually.
	if toAdd, ok := memo[v]; ok {
		return toAdd
	} else {
		// but just in case something retarded happened
		if v == 0 {
			memo[v] = 1
			return 1
		}

		// start memoizing
		// find the lowest possible combo, and start with it, because everything above it will rely upon it partially.
		lower := v - 1
		for lower >= 0 && joltage - s.adapters[lower] <= 3 {
			lower--
		}

		lower++ // get back on point, because we had to go one past to find the actual lowest.
		// process from lowest to highest
		var totalSolves int64
		for lower < v {
			totalSolves += s.RecursivePart2(lower, memo)

			lower++
		}

		// yes I spent 3 fucking hours not realizing I was missing this stupid goddamn if statement because I DID NOT READ THE ENTIRE PROBLEM CORRECTLY
		if s.adapters[v] <= 3 {
			totalSolves++
		}

		memo[v] = totalSolves
		return totalSolves
	}
}

func (s *Day10Solution) Part2() string {
	memo := map[int]int64{ 0:1 }
	count := s.RecursivePart2(len(s.adapters)-1, memo)

	return strconv.FormatInt(count, 10)
}
