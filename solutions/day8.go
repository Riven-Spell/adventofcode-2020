package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2020/util"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type Day8Solution struct {
	cpu gameConsoleCPU
}

func (s *Day8Solution) Prepare(input string) {
	instStrings := strings.Split(input, "\n")
	s.cpu.Rom = make([]gcCpuInstruction, len(instStrings))

	for k,v := range instStrings {
		instruction := instructions[v[:strings.Index(v, " ")]]
		count, err := strconv.ParseInt(strings.TrimPrefix(v[strings.Index(v," ")+1:], "+"), 10, 64)

		util.PanicIfErr(err)

		s.cpu.Rom[k] = gcCpuInstruction {
			instName: instruction,
			value:    count,
		}
	}
}

func (s *Day8Solution) Part1() string {
	s.cpu.Reset()
	seenInstructions := make(map[int64]bool)

	for _, ok := seenInstructions[s.cpu.ExecutionHead]; !ok; _, ok = seenInstructions[s.cpu.ExecutionHead] {
		fmt.Println(s.cpu.ExecutionHead, s.cpu.Accumulator)
		seenInstructions[s.cpu.ExecutionHead] = true

		s.cpu.Step()
	}

	return strconv.FormatInt(s.cpu.Accumulator, 10)
}

func (s *Day8Solution) Part2() string {
	s.cpu.Reset()
	registeredJumps := make(map[int]bool)

	for k,v := range s.cpu.Rom {
		if v.instName == JMP || v.instName == NOP {
			registeredJumps[k] = true
		}
	}

	var answer int64
	var wg sync.WaitGroup

	for k := range registeredJumps {
		if atomic.LoadInt64(&answer) != 0 {
			// early kill
			break
		}

		runner := s.cpu.Clone()

		if runner.Rom[k].instName == JMP {
			runner.Rom[k].instName = NOP
		} else {
			runner.Rom[k].instName = JMP
		}

		wg.Add(1)
		go runner.RunUntilTerminationOrForceHalt(&answer, &wg)
	}

	wg.Wait()

	return strconv.FormatInt(answer, 10)
}
