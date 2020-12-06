package tests

import (
	"github.com/Virepri/adventofcode-2020/solutions"
	"github.com/Virepri/adventofcode-2020/util"
	chk "gopkg.in/check.v1"
)

type Day5TestSuite struct{}
var _ = chk.Suite(&Day5TestSuite{})

func (s *Day5TestSuite) TestPassParse(c *chk.C) {
	solution := &solutions.Day5Solution{}

	c.Assert(solution.ParsePass("FBFBBFFRLR"), chk.Equals, util.Vector2D{44, 5})
	c.Assert(solution.ParsePass("BFFFBBFRRR"), chk.Equals, util.Vector2D{70, 7})
	c.Assert(solution.ParsePass("FFFBBBFRRR"), chk.Equals, util.Vector2D{14, 7})
	c.Assert(solution.ParsePass("BBFFBBFRLL"), chk.Equals, util.Vector2D{102,4})
}