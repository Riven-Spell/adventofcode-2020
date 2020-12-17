package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"strconv"
	"strings"
)

type Day17Solution struct {
	dimension map[util.Vector3D]bool
	initialVectorSize util.Vector3D
}

func (s *Day17Solution) Prepare(input string) {
	s.dimension = make(map[util.Vector3D]bool)

	rows := strings.Split(input, "\n")
	s.initialVectorSize.Y = int64(len(rows)) - 1
	s.initialVectorSize.X = int64(len(rows[0])) - 1

	for y, row := range rows {
		for x, col := range strings.Split(row, "") {
			if col == "#" {
				s.dimension[util.Vector3D{
					X: int64(x),
					Y: int64(y),
					Z: 0,
				}] = true
			}
		}
	}
}

func (s *Day17Solution) Part1() string {
	currentState := make(map[util.Vector3D]bool)
	workingState := make(map[util.Vector3D]bool)

	//min := util.Vector3D{}
	//max := s.initialVectorSize

	//adjustMapReading := func(vec util.Vector3D) {
	//	if vec.X > max.X {
	//		max.X = vec.X
	//	} else if vec.X < min.X {
	//		min.X = vec.X
	//	}
	//
	//	if vec.Y > max.Y {
	//		max.Y = vec.Y
	//	} else if vec.Y < min.Y {
	//		min.Y = vec.Y
	//	}
	//
	//	if vec.Z > max.Z {
	//		max.Z = vec.Z
	//	} else if vec.Z < min.Z {
	//		min.Z = vec.Z
	//	}
	//}
	//
	//printCurrentMap := func() {
	//	fmt.Println(min, max)
	//
	//	for z := min.Z; z <= max.Z; z ++ {
	//		fmt.Printf("z=%d\n", z)
	//		for y := min.Y; y <= max.Y; y ++ {
	//			for x := min.X; x <= max.X; x++ {
	//				state, ok := currentState[util.Vector3D{
	//					X: x,
	//					Y: y,
	//					Z: z,
	//				}]
	//
	//				if ok && !state {
	//					fmt.Print("!")
	//				} else {
	//					fmt.Print(util.TernaryString(ok, "#", "."))
	//				}
	//			}
	//			fmt.Println()
	//		}
	//		fmt.Println()
	//	}
	//}


	for loc,v := range s.dimension {
		currentState[loc] = v
	}

	//printCurrentMap()

	for i := 0; i < 6; i++ {
		pokedBits := make(map[util.Vector3D]int64)
		//min = util.Vector3D{math.MaxInt64, math.MaxInt64, math.MaxInt64}
		//max = util.Vector3D{math.MinInt64, math.MinInt64, math.MinInt64}

		// poke all adjacents
		for vec := range currentState {
			for _, a := range util.AdjacentVector3D {
				newVec := vec.Add(a)
				pokedBits[newVec]++
			}
		}

		// perform the rules based on pokes
		for vec := range currentState {
			pokes := pokedBits[vec]

			if pokes == 2 || pokes == 3 {
				//adjustMapReading(vec)
				workingState[vec] = true
			}
		}

		for vec, pokes := range pokedBits {
			if _, exists := currentState[vec]; !exists && pokes == 3 {
				//adjustMapReading(vec)
				workingState[vec] = true
			}
		}

		// clear the pokes
		pokedBits = make(map[util.Vector3D]int64)

		// swap the maps, copy current to working
		currentState = workingState
		workingState = make(map[util.Vector3D]bool)

		//fmt.Println("CYCLE:", i+1)
		//printCurrentMap()
	}

	return strconv.FormatInt(int64(len(currentState)), 10)
}

func (s *Day17Solution) Part2() string {
	// convert the initial state to a vector4d map
	currentState := map[util.Vector4D]bool{}
	workingState := map[util.Vector4D]bool{}

	for vec := range s.dimension {
		currentState[vec.To4D(0)] = true
	}

	for i := 0; i < 6; i++ {
		pokedBits := map[util.Vector4D]int64{}
		//min = util.Vector3D{math.MaxInt64, math.MaxInt64, math.MaxInt64}
		//max = util.Vector3D{math.MinInt64, math.MinInt64, math.MinInt64}

		// poke all adjacents
		for vec := range currentState {
			for _, a := range util.AdjacentVector4D {
				newVec := vec.Add(a)
				pokedBits[newVec]++
			}
		}

		// perform the rules based on pokes
		for vec := range currentState {
			pokes := pokedBits[vec]

			if pokes == 2 || pokes == 3 {
				//adjustMapReading(vec)
				workingState[vec] = true
			}
		}

		for vec, pokes := range pokedBits {
			if _, exists := currentState[vec]; !exists && pokes == 3 {
				//adjustMapReading(vec)
				workingState[vec] = true
			}
		}

		// swap the maps, copy current to working
		currentState = workingState
		workingState = make(map[util.Vector4D]bool)

		//fmt.Println("CYCLE:", i+1)
		//printCurrentMap()
	}

	return strconv.FormatInt(int64(len(currentState)), 10)
}