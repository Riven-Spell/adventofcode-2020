package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"strconv"
	"strings"
)

type Day16Solution struct {
	validityRanges map[string][][]int64
	myTicket []int64
	otherTickets [][]int64
}

func (s *Day16Solution) Prepare(input string) {
	sections := strings.Split(input, "\n\n")

	// parse validity ranges (sections[0])
	s.validityRanges = make(map[string][][]int64)
	validities := strings.Split(sections[0], "\n")
	for _,v := range validities {
		definition := strings.Split(v, ": ")
		ranges := strings.Split(definition[1], " or ")

		s.validityRanges[definition[0]] = make([][]int64, len(ranges))
		for k, r := range ranges {
			s.validityRanges[definition[0]][k] = util.ParseInts(r, "-")
		}
	}
	// parse my ticket (sections[1])
	mt := strings.Split(sections[1], "\n")[1]
	s.myTicket = util.ParseInts(mt, ",")
	// parse their tickets (sections[2])
	ot := strings.Split(sections[2], "\n")[1:] // cut off the nearby tickets line
	s.otherTickets = make([][]int64, len(ot))
	for k,v := range ot {
		s.otherTickets[k] = util.ParseInts(v, ",")
	}
}

func (s *Day16Solution) Part1() string {
	totalInvalid := uint64(0)

	// for each ticket, for each value, if it is not valid in at least one range, increment totalInvalid and skip to the next ticket.
	for _,v := range s.otherTickets {
		for _,val := range v {
			validAtLeastOnce := false

			for _, pol := range s.validityRanges {
				for _, r := range pol {
					if val >= r[0] && val <= r[1] {
						validAtLeastOnce = true
						goto skipCheckingValid
					}
				}
			}

			if !validAtLeastOnce {
				totalInvalid += uint64(util.IntAbs(val))
				goto skipTicket
			}

			skipCheckingValid:
		}

		skipTicket:
	}

	return strconv.FormatUint(totalInvalid, 10)
}



func (s *Day16Solution) Part2() string {
	// unionizers is a list, mapped to the individual values of a ticket.
	// this helps us gather the union of all possible fields
	unionizers := make([]util.Unionizer, len(s.otherTickets[0]))

	for _,v := range s.otherTickets {
		validPolicies := make([][]string, len(v))

		for valueID,val := range v {
			validProperties := make([]string, 0)

			for polName, pol := range s.validityRanges {
				for _, r := range pol {
					if val >= r[0] && val <= r[1] {
						validProperties = append(validProperties, polName)
					}
				}
			}

			if len(validProperties) == 0 {
				goto skipTicket
			}

			// set the valid properties
			validPolicies[valueID] = validProperties
		}

		// add valid properties to their respective unionizers
		for k,v := range validPolicies {
			unionizers[k].AddItems(v)
		}
	skipTicket:
	}

	// unionizer reduction funcs
	checkUnionizersEven := func() bool {
		for _,v := range unionizers {
			if v.Len() != 1 {
				return false
			}
		}

		return true
	}
	usedUnionizers := map[int]bool{}
	findLoneUnion := func() (string, int) {
		for k,v := range unionizers {
			if _, ok := usedUnionizers[k]; ok {
				continue // do not consider this unionizer because it has been found before
			}

			if v.Len() == 1 {
				usedUnionizers[k] = true
				return v.GetUnion()[0].(string), k
			}
		}

		panic("no lone unions left!")
	}

	//printUnionizerState := func() {
	//	for k,v := range unionizers {
	//		fmt.Printf("%d:", k)
	//		for _, val := range v.GetUnion() {
	//			fmt.Printf("%s, ", val)
	//		}
	//		fmt.Println()
	//	}
	//}

	// reduce the unions down to single entries
	for !checkUnionizersEven() {
		//printUnionizerState()
		firstLone, uIdx := findLoneUnion()
		//fmt.Println(firstLone, uIdx)

		for k,v := range unionizers {
			if k == uIdx || v.Len() == 1 {
				continue
			}

			v.RemoveItems([]string{firstLone})
		}
	}

	product := int64(1)
	for k,v := range unionizers {
		if strings.HasPrefix(v.GetUnion()[0].(string), "departure") {
			product *= s.myTicket[k]
		}
	}

	return strconv.FormatInt(product, 10)
}
