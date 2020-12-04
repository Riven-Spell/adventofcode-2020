package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"strconv"
	"strings"
	"sync"
)

type Day3Solution struct {
	scenery map[util.Vector2D]bool
	width, depth int
}

func (s *Day3Solution) Prepare(input string) {
	splits := strings.Split(input, "\n")

	s.scenery = make(map[util.Vector2D]bool)
	s.depth = len(splits)
	s.width = len(splits[0])

	for y, row := range splits {
		for x, char := range row {
			s.scenery[util.Vector2D{
				X: int64(x),
				Y: int64(y),
			}] = char == '#'
		}
	}
}

func (s *Day3Solution) Part1() string {
	activePoint := util.Vector2D{}
	collisions := int64(0)

	for activePoint.Y < int64(s.depth) {
		if s.scenery[activePoint] {
			collisions++
		}

		activePoint.Y++
		activePoint.X = (activePoint.X + 3) % int64(s.width)
	}

	return strconv.FormatInt(collisions, 10)
}

func (s *Day3Solution) Part2() string {
	slopes := []util.Vector2D{
		{1,1},
		{3,1},
		{5,1},
		{7,1},
		{1,2},
	}
	collisions := make([]int64, len(slopes))

	var wg sync.WaitGroup

	wg.Add(len(slopes))
	for n, v := range slopes {
		go func(slope util.Vector2D, cIndex int) {
			defer wg.Done()

			activePoint := util.Vector2D{}

			for activePoint.Y < int64(s.depth) {
				if s.scenery[activePoint] {
					collisions[cIndex]++
				}

				activePoint.Y += slope.Y
				activePoint.X = (activePoint.X + slope.X) % int64(s.width)
			}
		}(v, n)
	}

	wg.Wait()
	total := collisions[0]

	for _,v := range collisions[1:] {
		total *= v
	}

	return strconv.FormatInt(total, 10)
}