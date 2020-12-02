package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2020/util"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type PasswordPolicy struct {
	min, max int
	letter, password string // Technically this is bad data structuring but I don't care since there's one policy per password here
}

type Day2Solution struct{
	policies []PasswordPolicy
}

func (s *Day2Solution) Prepare(input string) {
	splits := strings.Split(input, "\n")
	s.policies = make([]PasswordPolicy, len(splits))

	for idx, pol := range splits {
		_, err := fmt.Sscanf(pol, "%d-%d %s %s", &s.policies[idx].min, &s.policies[idx].max, &s.policies[idx].letter, &s.policies[idx].password)

		s.policies[idx].letter = s.policies[idx].letter[:1]

		util.PanicIfErr(err)
	}
}

func (s *Day2Solution) Part1() string {
	var wg sync.WaitGroup
	var atomicValidity int64

	cpuCount := runtime.NumCPU()
	wg.Add(cpuCount)
	for i := 0; i < cpuCount; i++ {
		go func(cpu int) {
			defer wg.Done()

			for idx := cpu; idx < len(s.policies); idx += cpuCount {
				policy := s.policies[idx]
				count := strings.Count(policy.password, policy.letter)

				if count >= policy.min && count <= policy.max {
					atomic.AddInt64(&atomicValidity, 1)
				}
			}
		}(i)
	}

	wg.Wait()

	return strconv.FormatInt(atomic.LoadInt64(&atomicValidity), 10)
}


func (s *Day2Solution) Part2() string {
	var wg sync.WaitGroup
	var atomicValidity int64

	cpuCount := runtime.NumCPU()
	wg.Add(cpuCount)
	for i := 0; i < cpuCount; i++ {
		go func(cpu int) {
			defer wg.Done()

			for idx := cpu; idx < len(s.policies); idx += cpuCount {
				policy := s.policies[idx]

				if (policy.password[policy.min-1] == policy.letter[0]) != (policy.password[policy.max-1] == policy.letter[0]) {
					atomic.AddInt64(&atomicValidity, 1)
				}
			}
		}(i)
	}

	wg.Wait()

	return strconv.FormatInt(atomic.LoadInt64(&atomicValidity), 10)
}