package tests

import (
	"github.com/Virepri/adventofcode-2020/solutions"
	chk "gopkg.in/check.v1"
)

type Day8TestSuite struct{}
var _ = chk.Suite(&Day8TestSuite{})

func (*Day8TestSuite) TestDay8Part1(c *chk.C) {
	solution := &solutions.Day8Solution{}

	solution.Prepare(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)

	c.Assert(solution.Part1(), chk.Equals, "5")
}

func (*Day8TestSuite) TestDay8Part2(c *chk.C) {
	solution := &solutions.Day8Solution{}

	solution.Prepare(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)

	c.Assert(solution.Part2(), chk.Equals, "8")
}