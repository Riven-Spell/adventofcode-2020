package tests

import (
	"github.com/Virepri/adventofcode-2020/solutions"
	chk "gopkg.in/check.v1"
)

type Day10TestSuite struct{}
var _ = chk.Suite(&Day10TestSuite{})

func (*Day10TestSuite) TestDay10Part1(c *chk.C) {
	solution := &solutions.Day10Solution{}

	solution.Prepare(`16
10
15
5
1
11
7
19
6
12
4`)

	c.Assert(solution.Part1(), chk.Equals, "35")

	solution.Prepare(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)

	c.Assert(solution.Part1(), chk.Equals, "220")
}

func (*Day10TestSuite) TestDay10Part2(c *chk.C) {
	solution := &solutions.Day10Solution{}

	solution.Prepare(`16
10
15
5
1
11
7
19
6
12
4`)

	c.Assert(solution.Part2(), chk.Equals, "8")

	solution.Prepare(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)

	c.Assert(solution.Part2(), chk.Equals, "19208")
}