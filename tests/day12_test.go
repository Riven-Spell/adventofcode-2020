package tests

import (
	"github.com/Virepri/adventofcode-2020/solutions"
	chk "gopkg.in/check.v1"
)

type day12Suite struct{}
var _ = chk.Suite(&day12Suite{})

func (*day12Suite) TestDay12Part1(c *chk.C) {
	sol := &solutions.Day12Solution{}

	sol.Prepare(`F10
N3
F7
R90
F11`)

	c.Assert(sol.Part1(), chk.Equals, "25")
}

func (*day12Suite) TestDay12Part2(c *chk.C) {
	sol := &solutions.Day12Solution{}

	sol.Prepare(`F10
N3
F7
R90
F11`)

	c.Assert(sol.Part2(), chk.Equals, "286")
}