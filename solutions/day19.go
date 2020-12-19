package solutions

import (
	"github.com/Virepri/adventofcode-2020/util"
	"strconv"
	"strings"
)

type messageRule struct {
	subrules [][]int
	rule string // if == "", treat as subrules
}

func (mr messageRule) isValid(message string, rules []messageRule) []int {
	type internalState struct {
		ruleIDX int
		passed int
	}

	if mr.subrules != nil {
		validities := make([]int, 0)

		for _, ruleset := range mr.subrules {
			states := []internalState{{0, 0} }

			for len(states) > 0 {
				passed := states[0].passed
				rule := rules[ruleset[states[0].ruleIDX]]
				if passed >= len(message) {
					// cull this state since it's bad
					goto cull
				}

				if rule.rule != "" {
					if message[passed] != rule.rule[0] {
						//cull this one from the stack since it's bad
						goto cull
					}

					passed++

					if states[0].ruleIDX == len(ruleset) - 1 {
						// cull the state, add to validities
						validities = append(validities, passed)
					} else {
						states[0].passed = passed
						states[0].ruleIDX++
						goto noCull // keep handling the active state since this didn't split the stack at all.
					}
				} else {
					if counts := rule.isValid(message[passed:], rules); len(counts) > 0 {
						if states[0].ruleIDX == len(ruleset)-1 {
							// cull this state, add to validities
							for _, c := range counts {
								validities = append(validities, c+passed)
							}
						} else {
							// this rule is not the last, keep ticking.
							for _, c := range counts {
								states = append(states, internalState{
									ruleIDX: states[0].ruleIDX + 1,
									passed:  c + passed,
								})
							}
						}
					}
				}

				// cull the state
				cull:
				states = states[1:]
				noCull:
			}
		}

		return validities
	} else {
		//return message[0] == mr.rule[0], 1
		if message[0] == mr.rule[0] {
			return []int{1}
		}

		return []int{}
	}
}

type Day19Solution struct{
	rules []messageRule
	messages []string
}

func (s *Day19Solution) Prepare(input string) {
	parts := strings.Split(input, "\n\n")

	s.messages = strings.Split(parts[1], "\n")

	rules := strings.Split(parts[0], "\n")
	s.rules = make([]messageRule, len(rules))
	for _,v := range rules {
		splitIDX := strings.Index(v, ":")
		r := util.MustParseInt(v[:splitIDX])
		rule := messageRule{}

		if strings.HasSuffix(v[splitIDX+2:], "\"") {
			rule.rule = string(v[splitIDX+3])
		} else {
			rulesets := strings.Split(v[splitIDX+2:], " | ")
			rule.subrules = make([][]int, len(rulesets))

			for k, rs := range rulesets {
				rlist := strings.Split(rs, " ")

				rule.subrules[k] = make([]int, len(rlist))

				for idx,v := range rlist {
					rule.subrules[k][idx] = int(util.MustParseInt(v))
				}
			}
		}

		if len(s.rules) <= int(r) {
			s.rules = append(s.rules, make([]messageRule, int(r) - len(s.rules) + 1)...)
		}
		s.rules[r] = rule
	}
}

func (s *Day19Solution) Part1() string {
	count := 0
	for _,v := range s.messages {
		if validities := s.rules[0].isValid(v, s.rules); util.ArrayContains(validities, len(v)) {
			count++
		}
	}

	return strconv.FormatInt(int64(count), 10)
}

func (s *Day19Solution) Part2() string {
	count := 0

	tmpRules := make([]messageRule, len(s.rules))
	copy(tmpRules, s.rules)

	tmpRules[8] = messageRule{
		subrules: [][]int{ {42}, {42, 8} },
		rule:     "",
	}
	tmpRules[11] = messageRule{
		subrules: [][]int{ {42, 31}, {42, 11, 31} },
		rule:     "",
	}

	for _,v := range s.messages {
		if validities := tmpRules[0].isValid(v, tmpRules); util.ArrayContains(validities, len(v)) {
			count++
		}
	}

	return strconv.FormatInt(int64(count), 10)
}
