package tests

import (
	"github.com/Virepri/adventofcode-2020/solutions"
	chk "gopkg.in/check.v1"
)

type Day11Tests struct {}
var _ = chk.Suite(&Day11Tests{})

func (*Day11Tests) TestDay11Part1(c *chk.C) {
	solution := &solutions.Day11Solution{}

	solution.Prepare(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	c.Assert(solution.Part1(), chk.Equals, "37")
}

func (*Day11Tests) TestDay11Part2(c *chk.C) {
	solution := &solutions.Day11Solution{}

	solution.Prepare(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	c.Assert(solution.Part2(), chk.Equals, "26")
}
