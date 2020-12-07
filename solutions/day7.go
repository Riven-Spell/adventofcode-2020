package solutions

import (
	"strconv"
	"strings"
)

type Bag struct {
	Color string
	Contains map[string]int64
}

type Day7Solution struct {
	regulations map[string]Bag
}

func (s *Day7Solution) Prepare(input string) {
	s.regulations = make(map[string]Bag)
	regStrings := strings.Split(input, "\n")

	for _,v := range regStrings {
		bagColor := v[:strings.Index(v, " bags contain ")]

		a := v[len(bagColor + " bags contain "):]

		if a == "no other bags." {
			continue
		}

		bagsContain := strings.Split(a, ", ")

		b := Bag {
			Color: bagColor,
			Contains: make(map[string]int64),
		}

		for _, o := range bagsContain {
			firstSpace := strings.Index(o, " ")
			lastSpace := strings.LastIndex(o, " ")
			subBagQuantity, _ := strconv.ParseInt(o[:firstSpace], 10, 64)
			subBagColor := o[firstSpace+1:lastSpace]

			b.Contains[subBagColor] = subBagQuantity
		}

		s.regulations[bagColor] = b
	}
}

func (s *Day7Solution) Part1() string {
	var bagCount int64
	toContain := "shiny gold"

	for _,v := range s.regulations {
		if _, ok := v.Contains[toContain]; ok {
			bagCount++
		} else {
			bagQueue := make([]string, 0)

			for k := range v.Contains {
				bagQueue = append(bagQueue, k)
			}

			for len(bagQueue) > 0 {
				workItem := bagQueue[0]

				if _, ok := s.regulations[workItem].Contains[toContain]; ok {
					bagCount++
					goto breakBagQueue
				} else {
					for k, _ := range s.regulations[workItem].Contains {
						bagQueue = append(bagQueue, k)
					}
				}

				bagQueue = bagQueue[1:]
			}
		}

		breakBagQueue:
	}

	return strconv.FormatInt(bagCount, 10)
}

func (s *Day7Solution) Part2() string {
	type bagCount struct {
		Bag
		count int64
	}

	var total int64
	toCount := "shiny gold"
	workQueue := []bagCount{ {s.regulations[toCount], 1} }

	for len(workQueue) > 0 {
		bag := workQueue[0]

		for k,v := range bag.Contains {
			total += v * bag.count
			workQueue = append(workQueue, bagCount{ s.regulations[k], v * bag.count })
		}

		workQueue = workQueue[1:]
	}

	return strconv.FormatInt(total, 10)
}
