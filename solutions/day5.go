package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"math"
	"strconv"
	"strings"
)

type Bounds struct {
	top, bottom, left, right int64
}

type Day5Solution struct {
	Passes map[util.Vector2D]bool
}

func (s *Day5Solution) Prepare(input string) {
	s.Passes = make(map[util.Vector2D]bool)

 	for _,v := range strings.Split(input, "\n") {
		s.Passes[s.ParsePass(v)] = true
	}
}

func (s *Day5Solution) ParsePass(pass string) util.Vector2D {
	b := Bounds{
		top:    0,
		bottom: 127,
		left:   0,
		right:  7,
	}

	for _,c := range pass {
		switch c {
		case 'F':
			b.bottom = b.top + ((b.bottom - b.top) / 2)
		case 'B':
			b.top = b.top + ((b.bottom - b.top) / 2) + 1
		case 'R':
			b.left = b.left + ((b.right - b.left) / 2) + 1
		case 'L':
			b.right = b.left + ((b.right - b.left) / 2)
		}
	}

	return util.Vector2D{X: b.top, Y: b.right}
}

func (s *Day5Solution) Part1() string {
	var highest int64 = math.MinInt64

	for k, _ := range s.Passes {
		ID := (k.X * 8) + k.Y

		if ID > highest {
			highest = ID
		}
	}

	return strconv.FormatInt(highest, 10)
}
func (s *Day5Solution) Part2() string {
	var idToPass = make(map[int64]util.Vector2D)

	for k, _ := range s.Passes {
		ID := (k.X * 8) + k.Y

		idToPass[ID] = k
	}

	for k, _ := range idToPass {
		_, midOK := idToPass[k+1]
		_, rOK := idToPass[k+2]

		if !midOK && rOK {
			return strconv.FormatInt(k+1, 10)
		}
	}

	return "-1"
}