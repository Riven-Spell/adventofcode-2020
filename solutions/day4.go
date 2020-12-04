package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//type Passport struct {
//	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
//}

type Day4Solution struct{
	passports []map[string]string
}

func (s *Day4Solution) Prepare(input string) {
	splits := strings.Split(input, "\n\n")
	s.passports = make([]map[string]string, len(splits))

	for ppidx, pp := range splits {
		lines := strings.Split(pp, "\n")

		s.passports[ppidx] = make(map[string]string)

		for _, line := range lines {
			items := strings.Split(line, " ")

			for _, item := range items {
				idx := strings.Index(item, ":")

				if idx == -1 {
					panic("all items must be key:value")
				}

				s.passports[ppidx][item[:idx]] = item[idx+1:]
			}
		}
	}

	fmt.Println("")
}

func (s *Day4Solution) Part1() string {
	requiredFields := []string{ "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" }
	count := int64(0)

	for _, pp := range s.passports {
		good := int64(1)

		for _, field := range requiredFields {
			if _, ok := pp[field]; !ok {
				good = 0
				break
			}
		}

		count += good
	}

	return strconv.FormatInt(count, 10)
}

func (s *Day4Solution) Part2() string {
	count := int64(0)
	hclValidity := regexp.MustCompile("^#[0-9a-f]{6}$")
	eclValidity := map[string]bool{ "amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true }
	pidValidity := regexp.MustCompile("^[0-9]{9}$")

	for _, pp := range s.passports {
		good := int64(1)

		if byr, ok := pp["byr"]; ok {
			if i, err := strconv.ParseInt(byr, 10, 64); err != nil || i < 1920 || i > 2002 {
				goto skip
			}
		} else { goto skip }

		if iyr, ok := pp["iyr"]; ok {
			if i, err := strconv.ParseInt(iyr, 10, 64); err != nil || i < 2010 || i > 2020 {
				goto skip
			}
		} else { goto skip }

		if eyr, ok := pp["eyr"]; ok {
			if i, err := strconv.ParseInt(eyr, 10, 64); err != nil || i < 2020 || i > 2030 {
				goto skip
			}
		} else { goto skip }

		if hgt, ok := pp["hgt"]; ok {
			var max, min int64 = 0, 0

			if len(hgt) < 2 {
				goto skip
			}

			switch hgt[len(hgt)-2:] {
			case "cm":
				min, max = 150, 193
			case "in":
				min, max = 59, 76
			default:
				goto skip
			}

			if i, err := strconv.ParseInt(hgt[:len(hgt)-2], 10, 64); err != nil || i < min || i > max {
				goto skip
			}
		} else { goto skip }

		if hcl, ok := pp["hcl"]; ok {
			if !hclValidity.MatchString(hcl) {
				goto skip
			}
		} else { goto skip }

		if ecl, ok := pp["ecl"]; ok {
			if _, ok := eclValidity[ecl]; !ok {
				goto skip
			}
		} else { goto skip }

		if pid, ok := pp["pid"]; ok {
			if !pidValidity.MatchString(pid) {
				goto skip
			}
		} else { goto skip }

		goto jump

		skip:
			good = 0
		jump:

		count += good
	}

	return strconv.FormatInt(count, 10)
}
