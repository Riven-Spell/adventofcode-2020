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

func (mr messageRule) isValidP2(message string, rules []messageRule) []int {
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
				//all
				for k, idx := range ruleset[states[0].ruleIDX:] {
					if passed >= len(message) {
						// cull this state since it's bad
						states = states[1:]
						break
					}

					rule := rules[idx]

					if rule.rule != "" {
						if message[passed] != rule.rule[0] {
							//cull this one from the stack since it's bad
							states = states[1:]
							break
						}

						passed++

						if k + states[0].ruleIDX == len(ruleset) - 1 {
							// cull the state, add to validities
							validities = append(validities, passed)
							states = states[1:]
							break
						} else {
							states[0].passed = passed
							states[0].ruleIDX++
							break
						}
					} else {
						if counts := rule.isValidP2(message[passed:], rules); len(counts) > 0 {
							if k + states[0].ruleIDX == len(ruleset) - 1 {
								// cull this state, add to validities
								for _,c := range counts {
									validities = append(validities, c + passed)
								}
								states = states[1:]
								break
							} else {
								// this rule is not the last, keep ticking.
								for _,c := range counts {
									states = append(states, internalState{
										ruleIDX: states[0].ruleIDX + 1,
										passed:  c + passed,
									})
								}

								// cull the root state
								states = states[1:]
								break
							}
						} else {
							// cull this one from the stack since it's bad
							states = states[1:]
							break
						}
					}
				}
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

func (mr messageRule) isValid(message string, rules []messageRule) (valid bool, passed int) {
	if mr.subrules != nil {
		for _, ruleset := range mr.subrules {
			//each rule in the ruleset needs to come back true for its section
			passed := 0
			allValid := true
			for _, idx := range ruleset {
				if passed >= len(message) {
					// break, this is going too long.
					allValid = false
					break
				}

				rule := rules[idx]

				if rule.rule != "" {
					if message[passed] != rule.rule[0] {
						allValid = false
						break
					}

					passed++
				} else {
					if v, count := rule.isValid(message[passed:], rules); v {
						passed += count
					} else {
						allValid = false
						break
					}
				}
			}

			if allValid {
				return allValid, passed
			}
		}

		return false, 0
	} else {
		return message[0] == mr.rule[0], 1
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
		if ok, used := s.rules[0].isValid(v, s.rules); ok && used == len(v) {
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

	containsInt := func(i []int, x int) bool {
		for _,v := range i {
			if v == x {
				return true
			}
		}

		return false
	}

	for _,v := range s.messages {
		if validities := tmpRules[0].isValidP2(v, tmpRules); containsInt(validities, len(v)) {
			count++
		}
	}

	return strconv.FormatInt(int64(count), 10)
}
