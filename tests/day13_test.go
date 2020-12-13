package tests

import (
	"github.com/Virepri/adventofcode-2020/solutions"
	chk "gopkg.in/check.v1"
)

type Day13Tests struct{}
var _ = chk.Suite(&Day13Tests{})

func (*Day13Tests) TestPart1(c *chk.C) {
	s := &solutions.Day13Solution{}

	s.Prepare(`939
7,13,x,x,59,x,31,19`)

	c.Assert(s.Part1(), chk.Equals, "295")
}

func (*Day13Tests) TestPart2(c *chk.C) {
	s := &solutions.Day13Solution{}

	s.Prepare(`939
2,x,5`)

	c.Assert(s.Part2(), chk.Equals, "8")

	s.Prepare(`939
1,4,x,2`)

	c.Assert(s.Part2(), chk.Equals, "3")

	s.Prepare(`0
2,5,9`)

	c.Assert(s.Part2(), chk.Equals, "34")

	s.Prepare(`0
17,x,13,19`)

	c.Assert(s.Part2(), chk.Equals, "3417")

	s.Prepare(`939
7,13,x,x,59,x,31,19`)

	c.Assert(s.Part2(), chk.Equals, "1068781")

	s.Prepare(`0
1789,37,47,1889`)

	c.Assert(s.Part2(), chk.Equals, "1202161486")
}

