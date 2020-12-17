package tests

import (
	"github.com/Virepri/adventofcode-2020/util"
	chk "gopkg.in/check.v1"
)

type UnionizerTestSuite struct {}
var _ = chk.Suite(&UnionizerTestSuite{})

func (*UnionizerTestSuite) TestAddEmpty(c *chk.C) {
	u := util.Unionizer{}

	// add 2 and 5
	u.AddItems([]int{2,5})

	// check for both
	contains2 := false
	contains5 := false
	for _, v := range u.GetUnion() {
		if v.(int) == 2 {
			contains2 = true
		} else if v.(int) == 5 {
			contains5 = true
		}
	}
	c.Assert(contains2 && contains5, chk.Equals, true)
}

func (*UnionizerTestSuite) TestNoUnion(c *chk.C) {
	u := util.Unionizer{}

	// add 2 and 5
	u.AddItems([]int{2,5})
	// add a non-value
	u.AddItems([]int{0})
	// ensure that nothing is present
	c.Assert(u.Len(), chk.Equals, 0)
}

func (*UnionizerTestSuite) TestBasicUnion(c *chk.C) {
	u := util.Unionizer{}

	u.AddItems([]int{5,2})
	u.AddItems([]int{2,8})
	c.Assert(u.Len(), chk.Equals, 1)
}

func (*UnionizerTestSuite) TestRemoveItems(c *chk.C) {
	u := util.Unionizer{}

	u.AddItems([]int{0,1})
	c.Assert(u.Len(), chk.Equals, 2)
	u.RemoveItems([]int{1})
	c.Assert(u.Len(), chk.Equals, 1)
}