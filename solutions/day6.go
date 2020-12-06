package solutions

import (
	"strconv"
	"strings"
)

type Questionnaire struct {
	GroupAnswers map[rune]int
	GroupSize int
}

type Day6Solution struct{
	Questionnaires []Questionnaire
}

func (s *Day6Solution) Prepare(input string) {
	splits := strings.Split(input, "\n\n")
	s.Questionnaires = make([]Questionnaire, len(splits))

	for k,v := range splits {
		s.Questionnaires[k].GroupAnswers = make(map[rune]int)
		s.Questionnaires[k].GroupSize = 1

		for _,char := range v {
			if char == '\n' {
				s.Questionnaires[k].GroupSize++
				continue
			}

			s.Questionnaires[k].GroupAnswers[char]++
		}
	}
}

func (s *Day6Solution) Part1() string {
	sum := int64(0)

	for _,v := range s.Questionnaires {
		sum += int64(len(v.GroupAnswers))
	}

	return strconv.FormatInt(sum, 10)
}

func (s *Day6Solution) Part2() string {
	sum := int64(0)

	for _,v := range s.Questionnaires {
		for _, count := range v.GroupAnswers {
			if count == v.GroupSize {
				sum++
			}
		}
	}

	return strconv.FormatInt(sum, 10)
}