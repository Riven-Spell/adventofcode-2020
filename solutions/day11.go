package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"strconv"
	"strings"
)

type seatState uint8

const (
	FLOOR seatState = iota
	EMPTY
	TAKEN
)

type Day11Solution struct{
	maxX, maxY int64
	seatingMap map[util.Vector2D]seatState
}

func (s *Day11Solution) Prepare(input string) {
	s.seatingMap = make(map[util.Vector2D]seatState)

	rows := strings.Split(input, "\n")

	s.maxY = int64(len(rows))
	s.maxX = int64(len(rows[0]))

	for y,row := range rows {
		for x, char := range row {
			state := FLOOR

			if char == 'L' {
				state = EMPTY
			}

			s.seatingMap[util.Vector2D{X: int64(x), Y: int64(y)}] = state
		}
	}
}

func (s *Day11Solution) PrintMap(seating map[util.Vector2D] seatState) (output string) {
	location := util.Vector2D{}

	for ; location.Y < s.maxY ; location.Y++ {
		for ; location.X < s.maxX ; location.X++ {
			toWrite := "."
			state := seating[location]

			if state == TAKEN {
				toWrite = "#"
			} else if state == EMPTY {
				toWrite = "L"
			}

			output += toWrite
		}
		output += "\n"

		location.X = 0
	}

	return
}

func (s *Day11Solution) Part1() string {
	currentState := make(map[util.Vector2D]seatState) // note that the rule applies in atomic time-- Thus, new changes within the last tick are not considered in the rule
	workingMap := make(map[util.Vector2D]seatState)

	// copy the map
	for k, v := range s.seatingMap {
		workingMap[k] = v
		currentState[k] = v
	}

	for {
		for k, v := range currentState {
			if v == FLOOR {
				continue
			}

			adjacentFilled := 0
			for _, adj := range util.AdjacentVector2D {
				loc := k.Add(adj)

				s := currentState[loc]
				if s == TAKEN {
					adjacentFilled++
				}
			}

			if v == EMPTY && adjacentFilled == 0 {
				workingMap[k] = TAKEN
			} else if v == TAKEN && adjacentFilled >= 4 {
				workingMap[k] = EMPTY
			} else {
				workingMap[k] = v
			}
		}

		// swap them
		currentState, workingMap = workingMap, currentState
		//x := s.PrintMap(currentState)
		//fmt.Println(x + "\n")

		for k,v := range currentState {
			if workingMap[k] != v {
				goto skipBreak
			}
		}

		break
		skipBreak:
	}

	occupied := int64(0)
	for _,v := range currentState {
		if v == TAKEN {
			occupied++
		}
	}

	return strconv.FormatInt(occupied, 10)
}

func (s *Day11Solution) Part2() string {
	currentState := make(map[util.Vector2D]seatState) // note that the rule applies in atomic time-- Thus, new changes within the last tick are not considered in the rule
	workingMap := make(map[util.Vector2D]seatState)

	// copy the map
	for k, v := range s.seatingMap {
		workingMap[k] = v
		currentState[k] = v
	}

	for {
		for k, v := range currentState {
			if v == FLOOR {
				continue
			}

			//if k == (util.Vector2D{3, 0}) {
			//	x := s.PrintMap(currentState)
			//	fmt.Println(x)
			//}

			adjacentFilled := 0
			for _, adj := range util.AdjacentVector2D {
				toCheck := k.Add(adj)

				for toCheck.X >= 0 && toCheck.X < s.maxX && toCheck.Y >= 0 && toCheck.Y < s.maxY {
					s := currentState[toCheck]
					if s == TAKEN {
						adjacentFilled++
						toCheck = util.Vector2D{-1, -1}
					} else if s == EMPTY {
						toCheck = util.Vector2D{-1,-1}
					} else {
						toCheck = toCheck.Add(adj)
					}
				}
			}

			if v == EMPTY && adjacentFilled == 0 {
				workingMap[k] = TAKEN
			} else if v == TAKEN && adjacentFilled >= 5 {
				workingMap[k] = EMPTY
			} else {
				workingMap[k] = v
			}
		}

		// swap them
		currentState, workingMap = workingMap, currentState
		//x := s.PrintMap(currentState)
		//fmt.Println(x)
		//fmt.Println("--------------")

		for k,v := range currentState {
			if workingMap[k] != v {
				goto skipBreak
			}
		}

		break
	skipBreak:
	}

	occupied := int64(0)
	for _,v := range currentState {
		if v == TAKEN {
			occupied++
		}
	}

	return strconv.FormatInt(occupied, 10)
}
