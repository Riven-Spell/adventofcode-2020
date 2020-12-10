package solutions

/*
func (s *Day10Solution) RecursivePart2(v int, memo map[int]int64) int64 {
	joltage := s.adapters[v]

	//0 = 1 is already set, so there is no need to memoize it manually.
	if toAdd, ok := memo[v]; ok {
		return toAdd
	} else {
		// but just in case something retarded happened
		if v == 0 {
			memo[v] = 1
			return 1
		}

		// start memoizing
		// find the lowest possible combo, and start with it, because everything above it will rely upon it partially.
		lower := v - 1
		for lower >= 0 && joltage - s.adapters[lower] <= 3 {
			lower--
		}

		lower++ // get back on point, because we had to go one past to find the actual lowest.
		// process from lowest to highest
		var totalSolves int64
		for lower < v {
			totalSolves += s.RecursivePart2(lower, memo)

			lower++
		}

		memo[v] = totalSolves
		return totalSolves
	}
}
 */
