package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"strconv"
	"strings"
)

type direction uint8
type directionInstruction struct {
	dir direction
	count int64
}

const (
	FORWARD direction = iota
	LEFT
	RIGHT
	NORTH
	SOUTH
	EAST
	WEST
)

var directionParsingTable = map[rune]direction{ 'F':FORWARD, 'L':LEFT, 'R':RIGHT, 'N':NORTH, 'S':SOUTH, 'E': EAST, 'W':WEST }
var directionTranslation = map[direction]util.Vector2D{ NORTH: { 0, 1 }, SOUTH: { 0, -1 }, EAST: { 1, 0 }, WEST: { -1, 0 } }
var rotationTranslation = map[int]direction{ 0: WEST, 90: NORTH, 180: EAST, 270: SOUTH }

type Day12Solution struct {
	instructions []directionInstruction
}

func (s *Day12Solution) Prepare(input string) {
	toParse := strings.Split(input, "\n")
	s.instructions = make([]directionInstruction, len(toParse))
	
	for k,v := range toParse {
		s.instructions[k] = directionInstruction{
			dir:   directionParsingTable[rune(v[0])],
			count: util.MustParseInt(v[1:]),
		}
	}
}

func (s *Day12Solution) Part1() string {
	shipLocation := util.Vector2D{}
	shipRotation := 180

	for _,v := range s.instructions {
		switch v.dir {
		case FORWARD:
			shipLocation = shipLocation.Add(directionTranslation[rotationTranslation[shipRotation]].Mul(v.count))
		case LEFT:
			shipRotation -= int(v.count)
			shipRotation = util.EnsureRotation360(shipRotation)
		case RIGHT:
			shipRotation += int(v.count)
			shipRotation = util.EnsureRotation360(shipRotation)
		default:
			shipLocation = shipLocation.Add(directionTranslation[v.dir].Mul(v.count))
		}
	}

	return strconv.FormatInt(util.Manhattan(util.Vector2D{}, shipLocation), 10)
}

func (s *Day12Solution) Part2() string {
	shipLocation := util.Vector2D{}
	waypointLocation := directionTranslation[EAST].Mul(10).Add(directionTranslation[NORTH])

	for _,v := range s.instructions {
		switch v.dir {
		case FORWARD:
			shipLocation = shipLocation.Add(waypointLocation.Mul(v.count))
		case LEFT:
			waypointLocation = waypointLocation.Rot(float64(+v.count))
		case RIGHT:
			waypointLocation = waypointLocation.Rot(float64(-v.count))
		default:
			waypointLocation = waypointLocation.Add(directionTranslation[v.dir].Mul(v.count))
		}
	}

	return strconv.FormatInt(util.Manhattan(util.Vector2D{}, shipLocation), 10)
}
