package solutions

import "github.com/Virepri/adventofcode-2020/inputs"

// Solution is an interface that exposes a simple structure:
// The runner should call Prepare() on it to prepare the input
// Then call the part 1 and part 2 functions if wanted.
type Solution interface {
	Prepare(input string)
	Part1() string
	Part2() string
}

var RegisteredDays = []struct{
	Solution        Solution // sample solution can be found in ./sampleday.go
	StringInput     *string // inputs should exist in ../inputs and be a single var in a single file. These are just default inputs.
	ExpectedOutputs []string // these determine pass/failure in case I come back to try and optimize a solution and fuck it up.
}{
	{
		Solution:        &Day1Solution{},
		StringInput:     &inputs.Day1Input,
		ExpectedOutputs: []string{"969024", "230057040"},
	},
	{
		Solution:        &Day2Solution{},
		StringInput:     &inputs.Day2Input,
		ExpectedOutputs: []string{"424", "747"},
	},
	{
		Solution:        &Day3Solution{},
		StringInput:     &inputs.Day3Input,
		ExpectedOutputs: []string{"272", "3898725600"},
	},
	{
		Solution:        &Day4Solution{},
		StringInput:     &inputs.Day4Input,
		ExpectedOutputs: []string{"235", "194"},
	},
	{
		Solution: &Day5Solution{},
		StringInput: &inputs.Day5Input,
		ExpectedOutputs: []string{"888", "522"},
	},
	{
		Solution: &Day6Solution{},
		StringInput: &inputs.Day6Input,
		ExpectedOutputs: []string{"6930", "3585"},
	},
	{
		Solution: &Day7Solution{},
		StringInput: &inputs.Day7Input,
		ExpectedOutputs: []string{"169", "82372"},
	},
	{
		Solution: &Day8Solution{},
		StringInput: &inputs.Day8Input,
		ExpectedOutputs: []string{"1675", "1532"},
	},
}
