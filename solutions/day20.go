package solutions

import (
	"fmt"
	"github.com/Virepri/adventofcode-2020/util"
	"math"
	"strings"
)

type ImageAccessMethods struct {
	flipX, flipY bool
	rotations int
}

var rootFlips = []ImageAccessMethods {
	{},
	{true, false, 0},
	{false, true, 0},
	{true, true, 0},
}

var validRotations = func () []ImageAccessMethods{
	rootMethods := make([]ImageAccessMethods, len(rootFlips))
	copy(rootMethods, rootFlips)

	tmpMethods := make([]ImageAccessMethods, len(rootMethods) * 6)
	for k,v := range rootMethods {
		for r := 1; r <= 3; r++ {
			if r == 0 {
				continue
			}

			tmpMethods[(k*3)+r+3-1] = ImageAccessMethods{
				flipX:     v.flipX,
				flipY:     v.flipY,
				rotations: r,
			}
		}
	}
	rootMethods = append(rootMethods, tmpMethods...)

	return rootMethods
}()

type Image struct {
	size util.Vector2D
	bits [][]bool
}

func (i *Image) toString(am ImageAccessMethods) string {
	output := ""
	for y := int64(0); y < i.size.Y; y++ {
		for x := int64(0); x < i.size.X; x++ {
			output += util.TernaryString(i.GetPoint(util.Vector2D{X: x, Y: y}, am), "#", ".")
		}
		output += "\n"
	}
	return output
}

func (i *Image) find(d util.Vector2D, dest Image, iIAM ImageAccessMethods) (matches bool) {
	inc := util.Vector2D{}
	for ; inc.Y < dest.size.Y; inc.Y++ {
		for ; inc.X < dest.size.X; inc.X++ {
			thisImg := i.GetPoint(d.Add(inc), iIAM)
			thatImg := dest.bits[inc.Y][inc.X]

			if thatImg {
				if thisImg != thatImg {
					return false
				}
			}
		}
		inc.X = 0
	}

	return true
}

func (i *Image) clearBits(d util.Vector2D, dest Image, iIAM ImageAccessMethods) {
	inc := util.Vector2D{}
	for ; inc.Y < dest.size.Y; inc.Y++ {
		for ; inc.X < dest.size.X; inc.X++ {
			thatImg := dest.bits[inc.Y][inc.X]

			if thatImg {
				p := i.getAccessPoint(d.Add(inc), iIAM)

				i.bits[p.Y][p.X] = false
				//thisImg := i.GetPoint(d.Add(inc), iIAM)
			}
		}
		inc.X = 0
	}
}

func (i *Image) getAccessPoint(d util.Vector2D, am ImageAccessMethods) util.Vector2D {
	if am == (ImageAccessMethods{}) {
		return d
	}

	root := util.Vector2D{}
	for am.rotations < 0 {
		am.rotations += 4
	}
	switch am.rotations % 4 {
	case 1:
		am.flipX = !am.flipX
		d.X, d.Y = d.Y, d.X
	case 2:
		am.flipX = !am.flipX
		am.flipY = !am.flipY
	case 3:
		am.flipY = !am.flipY
		d.X, d.Y = d.Y, d.X
	}

	if am.flipX {
		d.X = -d.X
		root.X = i.size.X - 1
	}

	if am.flipY {
		d.Y = -d.Y
		root.Y = i.size.Y - 1
	}

	// subtract one from the size so that we are accessing the actual end
	return root.Add(d)
}

func (i *Image) GetPoint(d util.Vector2D, am ImageAccessMethods) bool {
	accessPoint := i.getAccessPoint(d, am)

	return i.bits[accessPoint.Y][accessPoint.X]
}

func (i *Image) FinalLock (left, upper Image, lRot, uRot ImageAccessMethods) ImageAccessMethods {
	for _,v := range validRotations {
		if i.CompareSides(v, left, lRot, 3) && i.CompareSides(v, upper, uRot, 0) {
			return v
		}
	}

	panic("no final lock")
}

func (i *Image) LockRotation (lower, right Image) (swapped bool, midRot, lowerRot, rightRot ImageAccessMethods) {
	retry:
	for _,v := range validRotations {
		for _, lRot := range validRotations {
			for _, rRot := range validRotations {
				// compare right side
				if i.CompareSides(v, lower, lRot, 2) && i.CompareSides(v, right, rRot, 1) {
					//fmt.Println(swapped, v, lRot, rRot)
					midRot, lowerRot, rightRot = v, lRot, rRot
					//fmt.Println(lower.toString(lRot))
					//fmt.Println(i.toString(v))
					//fmt.Println(right.toString(rRot))
					return
				}
			}
		}
	}

	if !swapped {
		lower, right = right, lower
		swapped = true
		goto retry
	}

	panic("no valid rotations to lock into!")
}

func (i *Image) CompareSides(iam ImageAccessMethods, i2 Image, i2am ImageAccessMethods, iSide int) bool {
	i2Side := (iSide + 2) % 4
	iRoot, iInc, iCount := i.GetSideTestingPoints(iSide)
	i2Root, i2Inc, i2Count := i2.GetSideTestingPoints(i2Side)

	if i2Count != iCount {
		panic("sides must be same length")
	}

	for inc := int64(0); inc < iCount; inc++ {
		iIDX := iRoot.Add(iInc.Mul(inc))
		i2DX := i2Root.Add(i2Inc.Mul(inc))

		if i.GetPoint(iIDX, iam) != i2.GetPoint(i2DX, i2am) {
			return false
		}
	}

	return true
}

// every pair of two methods is a valid combo.
func (i *Image) CompatibleRotations(i2 Image) []ImageAccessMethods {
	combos := make([]ImageAccessMethods, 0)

	for _, iam := range validRotations {
		for _, i2am := range validRotations {
			if i.CompareSides(iam, i2, i2am, 0) {
				combos = append(combos, iam, i2am)
			}
		}
	}

	return combos
}

func (i *Image) GetSideTestingPoints(side int) (root util.Vector2D, incrementor util.Vector2D, iCount int64) {
	switch side {
	case 0: // up
		return util.Vector2D{}, util.Vector2D{X: 1}, i.size.X
	case 1: // right
		return util.Vector2D{X: i.size.X - 1}, util.Vector2D{Y: 1}, i.size.Y
	case 2: // bottom
		return util.Vector2D{Y: i.size.Y - 1}, util.Vector2D{X: 1}, i.size.X
	case 3: // left
		return util.Vector2D{}, util.Vector2D{Y: 1}, i.size.Y
	}
	panic("no such direction")
}

type Day20Solution struct{
	images map[int]Image
	seaMonster Image
}

func (s *Day20Solution) Prepare(input string) {
	imgSplits := strings.Split(input, "\n\n")

	s.images = make(map[int]Image)

	for _,v := range imgSplits {
		lines := strings.Split(v, "\n")

		var imgIDX int
		_, err := fmt.Sscanf(lines[0], "Tile %d:", &imgIDX)
		util.PanicIfErr(err)

		img := Image{
			size: util.Vector2D{
				X: int64(len(lines[1])),
				Y: int64(len(lines) - 1), // Y has a title line to be subtracted
			},
			bits: make([][]bool, len(lines) - 1),
		}

		for y, row := range lines[1:] {
			img.bits[y] = make([]bool, len(row))
			for x, char := range row {
				img.bits[y][x] = char == '#'
			}
		}

		s.images[imgIDX] = img
	}

	seaMonster := `..................#.
#....##....##....###
.#..#..#..#..#..#...`

	lines := strings.Split(seaMonster, "\n")

	smIMG := Image{
		size: util.Vector2D{
			X: int64(len(lines[1])),
			Y: int64(len(lines)), // Y has a title line to be subtracted
		},
		bits: make([][]bool, len(lines)),
	}

	for y, row := range lines {
		smIMG.bits[y] = make([]bool, len(row))
		for x, char := range row {
			smIMG.bits[y][x] = char == '#'
		}
	}

	s.seaMonster = smIMG
}

func (s *Day20Solution) Part1() string {
	multiple := 1

	for k, v := range s.images {
			matches := make([]int, 0)
			for k2, v2 := range s.images {
				if k == k2 {
					continue // don't process a combo of the same ones
				}

				// tests all valid rotations and flips for valid connections
				rots := v.CompatibleRotations(v2)
				if len(rots) != 0 {
					matches = append(matches, k2)
				}
			}

			if len(matches) == 2 {
				multiple *= k
			}
	}

	return fmt.Sprint(multiple)
}

func (s *Day20Solution) Part2() string {
	connections := map[int][]int{}
	corners := []int{}

	for k, v := range s.images {
		matches := make([]int, 0)
		for k2, v2 := range s.images {
			if k == k2 {
				continue // don't process a combo of the same ones
			}

			// tests all valid rotations and flips for valid connections
			rots := v.CompatibleRotations(v2)
			if len(rots) != 0 {
				matches = append(matches, k2)
			}
		}

		connections[k] = matches
		if len(matches) == 2 {
			corners = append(corners, k)
		}
	}

	sideLength := int(math.Round(math.Sqrt(float64(len(s.images)))))
	images := make([][]util.Unionizer, sideLength)
	rotations := make([][]ImageAccessMethods, sideLength)

	for k := range images {
		rotations[k] = make([]ImageAccessMethods, sideLength)
		images[k] = make([]util.Unionizer, sideLength)
	}

	// treat the first found corner as the corner of the image, and grow out from there.
	images[0][0].AddItems([]int{corners[0]})
	used := []int{corners[0]}

	cRoot := util.Vector2D{}

	// expects the lower image of the trio
	getImageTrio := func(y,x int64) (lower, mid, right Image, lIDX, rIDX int) {
		//return s.images[images[y][x].GetUnion()[0].(int)]

		// mid is guaranteed locked in sans rotation
		mid = s.images[images[y-1][x].GetUnion()[0].(int)]
		// the union will only return two acceptable items
		runion := images[y-1][x+1]
		lunion := images[y][x]
		union := lunion.JoinUnions(runion)
		u := union.GetUnion()

		if len(u) > 2 {
			panic("union was larger than two")
		} else if len(u) == 1 {
			if runion.Len() == 1 {
				images[y][x].RemoveItems(runion.GetUnion())
			} else {
				images[y-1][x+1].RemoveItems(lunion.GetUnion())
			}

			rIDX = runion.GetUnion()[0].(int)
			lIDX = lunion.GetUnion()[0].(int)
			lower = s.images[lIDX]
			right = s.images[rIDX]

			return
		} else if len(u) == 0 {
			if runion.Len() == 1 && lunion.Len() == 1 {
				rIDX = runion.GetUnion()[0].(int)
				lIDX = lunion.GetUnion()[0].(int)
				lower = s.images[lIDX]
				right = s.images[rIDX]

				return
			}

			panic("union had nothing to offer")
		}

		lIDX = u[0].(int)
		rIDX = u[1].(int)
		lower = s.images[lIDX]
		right = s.images[rIDX]

		return
	}

	for cRoot.X <= int64(sideLength - 1) && cRoot.Y <= int64(sideLength - 1) {
		tmpRoot := cRoot

		//for tmpRoot.Y >= 0 && tmpRoot.X < int64(sideLength) {
		//	//rootConn := func() []int {
		//	//	out := make([]int, 0)
		//	//
		//	//	images[tmpRoot.Y][tmpRoot.X].ForEach(func(i interface{}) bool {
		//	//		out = append(out, connections[i.(int)]...)
		//	//		return true
		//	//	})
		//	//
		//	//	return out
		//	//}()

		for tmpRoot.Y >= 0 && tmpRoot.X < int64(sideLength - 1) {
			// lock in rotations for this first bit
			if tmpRoot != (util.Vector2D{}) && tmpRoot.Y != 0 {
				low, mid, right, lidx, ridx := getImageTrio(tmpRoot.Y, tmpRoot.X)
				swapped, midRot, lowerRot, rightRot := mid.LockRotation(low, right)

				if swapped {
					lidx, ridx = ridx, lidx
					lowerRot, rightRot = rightRot, lowerRot
				}

				rotations[tmpRoot.Y][tmpRoot.X] = lowerRot
				rotations[tmpRoot.Y-1][tmpRoot.X] = midRot
				rotations[tmpRoot.Y-1][tmpRoot.X+1] = rightRot

				images[tmpRoot.Y][tmpRoot.X].RemoveItems([]int{ridx})
				images[tmpRoot.Y-1][tmpRoot.X+1].RemoveItems([]int{lidx})
			}

			// then forward suspicions
			used = append(used, images[tmpRoot.Y][tmpRoot.X].GetUnion()[0].(int))
			if tmpRoot.Y < int64(sideLength) - 1 {
				images[tmpRoot.Y+1][tmpRoot.X].AddItems(connections[images[tmpRoot.Y][tmpRoot.X].GetUnion()[0].(int)])
				images[tmpRoot.Y+1][tmpRoot.X].RemoveItems(used)
			}
			if tmpRoot.X < int64(sideLength) - 1 {
				images[tmpRoot.Y][tmpRoot.X+1].AddItems(connections[images[tmpRoot.Y][tmpRoot.X].GetUnion()[0].(int)])
				images[tmpRoot.Y][tmpRoot.X+1].RemoveItems(used)
			}
			//images[tmpRoot.Y][tmpRoot.X].ForEach(func(i interface{}) bool {
			//	c := connections[i.(int)]
			//
			//	u.AddItems(c)
			//	u.RemoveItems(used)
			//})

			tmpRoot = tmpRoot.Add(util.Vector2D{1,-1})
		}

		//	// increment tmproot
		//	tmpRoot = tmpRoot.Add(util.Vector2D{1, -1})
		//}

		if cRoot == (util.Vector2D{int64(sideLength - 1), int64(sideLength - 1)}) {
			m := s.images[images[cRoot.Y][cRoot.X].GetUnion()[0].(int)]
			l := s.images[images[cRoot.Y][cRoot.X-1].GetUnion()[0].(int)]
			u := s.images[images[cRoot.Y-1][cRoot.X].GetUnion()[0].(int)]
			rotations[cRoot.Y][cRoot.X] = m.FinalLock(l, u, rotations[cRoot.Y][cRoot.X-1], rotations[cRoot.Y-1][cRoot.X])
		}

		// increment the current root
		if cRoot.Y == int64(sideLength - 1) {
			cRoot.X++
		} else {
			cRoot.Y++
		}
	}

	// print out arrangement
	//for y,v := range images {
	//	rows := make([][]string, len(v))
	//	for x, u := range v {
	//		img := s.images[u.GetUnion()[0].(int)]
	//		rows[x] = strings.Split(img.toString(rotations[y][x]), "\n")
	//	}
	//
	//	for row := range rows[0] {
	//		for _, v := range rows {
	//			fmt.Print(v[row] + " ")
	//		}
	//		fmt.Println()
	//	}
	//}

	// fucking 400 lines in and this is only the first fucking HALF of the goddamned solution.
	getImage := func(y, x int) (Image, ImageAccessMethods) {
		return s.images[images[y][x].GetUnion()[0].(int)], rotations[y][x]
	}

	// time to compile the images together
	compiledImage := Image{
		size: util.Vector2D{ int64(sideLength * 8), int64(sideLength * 8) },
		bits: make([][]bool, sideLength * 8),
	}

	for y := 0; y < len(compiledImage.bits); y += 8 {
		for tmpY := y; tmpY < y + 8; tmpY++ {
			compiledImage.bits[tmpY] = make([]bool, sideLength*8)
		}

		for imgX := 0; imgX < sideLength; imgX++ {
			x, xiam := getImage(y / 8, imgX)

			for internalY,row := range x.bits[1:len(x.bits)-1] {
				for internalX,_ := range row[1:len(row)-1] {
					compiledImage.bits[y + internalY][(imgX * 8) + internalX] = x.GetPoint(util.Vector2D{int64(internalX + 1), int64(internalY + 1)}, xiam)
				}
			}
		}
	}

	sm := s.seaMonster

	total := 0
	for _,v := range validRotations {
		matched := false
		root := util.Vector2D{}
		for ; root.Y < compiledImage.size.Y-sm.size.Y; root.Y++ {
			for ; root.X < compiledImage.size.X-sm.size.X; root.X++ {
				tmpMatched := compiledImage.find(root, sm, v)

				if tmpMatched {
					matched = true
					compiledImage.clearBits(root, sm, v)
				}
			}
			root.X = 0
		}

		if matched {
			for Y := int64(0); Y < compiledImage.size.Y; Y++ {
				for X := int64(0); X < compiledImage.size.X; X++ {
					if compiledImage.bits[Y][X] {
						total++
					}
				}
			}
			break
		}
	}

	return fmt.Sprint(total)
}

