package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	command string
	locator int
	value string
}

var setInst = regexp.MustCompile("")

type Day14Solution struct{
	instructions []instruction
}

func (s *Day14Solution) Prepare(input string) {
	commands := strings.Split(input, "\n")
	s.instructions = make([]instruction, len(commands))

	for k,v := range commands {
		cmdParts := strings.Split(v, " = ")

		value := cmdParts[1]
		cmdName := cmdParts[0]
		locator := int64(-1)
		if sIndex := strings.Index(cmdParts[0], "["); sIndex != -1 {
			cmdName = cmdParts[0][:sIndex]
			locator = util.MustParseInt(cmdParts[0][sIndex+1:len(cmdParts[0])-1])
		}

		s.instructions[k] = instruction{
			command: cmdName,
			locator: int(locator),
			value:   value,
		}
	}
}

func (s *Day14Solution) Part1() string {
	mem := map[int]uint64{}

	var newBits uint64
	var bitMask uint64
	for _,v := range s.instructions {
		switch v.command {
		case "mask":
			newBits = 0
			bitMask = 0
			for c := 0; c < len(v.value); c++ {
				switch v.value[c] {
				case 'X':
					bitMask <<= 1
					newBits <<= 1
				case '1':
					bitMask = (bitMask << 1) | 1
					newBits = (newBits << 1) | 1
				case '0':
					bitMask = (bitMask << 1) | 1
					newBits <<= 1
				}
			}
		case "mem":
			mem[v.locator] = (util.MustParseUint(v.value) & (^bitMask)) | (newBits & bitMask)
		}
	}

	sum := uint64(0)
	for _,v := range mem {
		sum += uint64(v)
	}

	return strconv.FormatUint(sum, 10)
}

func (s *Day14Solution) GetBitmaskPermutations(newBits, bitMask, dest uint64) []uint64 {
	base := dest & bitMask // clear the floating bits
	base = base | newBits // overwrite ones


	toClear := uint64(math.MaxUint64)
	toClear <<= 36

	out := []uint64{base}
	floaters := ^bitMask & ^toClear
	pushed := 0
	for floaters > 0 {
		if floaters & 1 == 1 {
			toWrite := uint64(1) << pushed
			toAppend := make([]uint64, len(out))

			for k,v := range out {
				toAppend[k] = v | toWrite
			}

			out = append(out, toAppend...)
		}

		// push floaters forward.
		floaters >>= 1
		pushed++
	}

	return out
}

func (s *Day14Solution) Part2() string {
	mem := map[uint64]uint64{}

	var newBits uint64
	var bitMask uint64
	for _,v := range s.instructions {
		switch v.command {
		case "mask":
			newBits = 0
			bitMask = 0
			for c := 0; c < len(v.value); c++ {
				switch v.value[c] {
				case 'X':
					bitMask <<= 1
					newBits <<= 1
				case '1':
					bitMask = (bitMask << 1) | 1
					newBits = (newBits << 1) | 1
				case '0':
					bitMask = (bitMask << 1) | 1
					newBits <<= 1
				}
			}
		case "mem":
			perms := s.GetBitmaskPermutations(newBits, bitMask, uint64(v.locator))

			for _,newLoc := range perms {
				mem[newLoc] = util.MustParseUint(v.value)
			}
		}
	}

	sum := uint64(0)
	for _,v := range mem {
		sum += v
	}

	return strconv.FormatUint(sum, 10)
}
