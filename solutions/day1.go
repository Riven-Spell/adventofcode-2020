package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"strconv"
	"sync/atomic"
)

type Day1Solution struct {
	input []int64
}

func (s *Day1Solution) Prepare(input string) {
	s.input = util.ParseInts(input, "\n")
}

func (s *Day1Solution) Part1() string {
	answers := make(chan []int, 1)
	var atomicDone int32 = 0 // there are multiple answers on a very technical basis, and intermittently this will deadlock if we don't atomically handle that.

	findAns := func(index int) {
		for idx, i := range s.input {
			if atomic.LoadInt32(&atomicDone) != 0 {
				return
			}

			if s.input[index]+i == 2020 {
				atomic.AddInt32(&atomicDone, 1)
				answers <- []int{index, idx}
				return
			}
		}
	}

	for idx := range s.input {
		// there's an opportunity of cracking this answer brute-force really early
		if atomic.LoadInt32(&atomicDone) != 0 {
			break
		}

		go findAns(idx)
	}

	output := <- answers

	return strconv.FormatInt(s.input[output[0]] * s.input[output[1]], 10)
}

func (s *Day1Solution) Part2() string {
	answers := make(chan []int, len(s.input))
	var atomicDone int32 = 0 // there are multiple answers on a very technical basis, and intermittently this will deadlock if we don't atomically handle that.

	findAns := func(index int) {

		findAns2 := func(index2 int) {
			for idx, i := range s.input {
				if atomic.LoadInt32(&atomicDone) != 0 {
					return
				}

				if s.input[index] + s.input[index2] + i == 2020 {
					atomic.AddInt32(&atomicDone, 1)
					answers <- []int{index, index2, idx}
					return
				}
			}
		}

		for idx := range s.input {
			if atomic.LoadInt32(&atomicDone) != 0 {
				return
			}

			go findAns2(idx)
		}
	}

	for idx := range s.input {
		// there's an opportunity of cracking this answer brute-force really early
		// as in this is varying between 10ms and literally a couple hundred microseconds.
		if atomic.LoadInt32(&atomicDone) != 0 {
			break
		}

		go findAns(idx)
	}

	output := <- answers

	return strconv.FormatInt(s.input[output[0]] * s.input[output[1]] * s.input[output[2]], 10)
}