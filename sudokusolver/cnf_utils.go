package sudokusolver

func cnfAtLeast1(c *CNF, lits []int) [][]int {
	for _, lit := range lits {
		if exists := c.lookup(lit); exists {
			return [][]int{}
		}
	}

	return [][]int{lits}
}

func cnfAtMost1(c *CNF, lits []int) [][]int {
	filteredLits := []int{}
	for _, lit := range lits {
		if exists := c.lookup(-lit); !exists {
			filteredLits = append(filteredLits, lit)
		}
	}

	// return cnfAtMost1Pairwise(c, filteredLits)
	// return cnfAtMost1Commander(c, filteredLits)
	return cnfAtMost1Bimander(c, filteredLits)
}

func cnfExactly1(c *CNF, lits []int) [][]int {
	result := make([][]int, 0, 1+len(lits)*len(lits)/2)

	result = append(result, cnfAtLeast1(c, lits)...)
	result = append(result, cnfAtMost1(c, lits)...)

	return result
}

func cnfAtMost1Pairwise(c *CNF, lits []int) [][]int {
	result := make([][]int, 0, len(lits)*len(lits)/2)
	for i := 0; i < len(lits); i++ {
		for j := i + 1; j < len(lits); j++ {
			// each pair can't be true at the same time
			result = append(result, cnfAtLeast1(c, []int{-lits[i], -lits[j]})...)
		}
	}
	return result
}

// Will Klieber and Gihwon Kwon. 2007.
// Efficient CNF Encoding for Selecting 1 from N Objects.
func cnfAtMost1Commander(c *CNF, lits []int) [][]int {
	if len(lits) <= 3 {
		return cnfAtMost1Pairwise(c, lits)
	}

	factor := 3
	m := (len(lits) + factor - 1) / factor
	result := make([][]int, 0, len(lits)*len(lits)/2)

	groups := make([][]int, m)
	for i := 0; i < m; i++ {
		groups[i] = lits[factor*i : min(factor*(i+1), len(lits))]
		// 1. At most one variable in a group can be true
		result = append(result, cnfAtMost1Pairwise(c, groups[i])...)
	}

	commanders := makeRange(c.nbVar+1, c.nbVar+m)
	c.nbVar += m

	//  2. If the commander variable of a group is false,
	// then none of the variables in the group can be true
	for i, commander := range commanders {
		for _, lit := range groups[i] {
			// -commander -> -lit
			result = append(result, cnfAtLeast1(c, []int{commander, -lit})...)
		}
	}

	// 3. Exactly one of the commander variables is true
	result = append(result, cnfAtLeast1(c, commanders)...)
	result = append(result, cnfAtMost1Commander(c, commanders)...)

	return result
}

// Nguyen, Van-Hau, and Son T. Mai. 2015.
// A new method to encode the at-most-one constraint into SAT.
func cnfAtMost1Bimander(c *CNF, lits []int) [][]int {
	factor := 2
	m := (len(lits) + factor - 1) / factor
	result := make([][]int, 0, len(lits)*len(lits)/2)

	groups := make([][]int, m)
	for i := 0; i < m; i++ {
		groups[i] = lits[factor*i : min(factor*(i+1), len(lits))]
		result = append(result, cnfAtMost1Pairwise(c, groups[i])...)
	}

	binLength := getBinLength(m)
	auxVars := makeRange(c.nbVar+1, c.nbVar+binLength)
	c.nbVar += binLength

	for i := 0; i < m; i++ {
		for _, lit := range groups[i] {
			for k, aux := range auxVars {
				commanderLit := aux
				if (i & (1 << k)) == 0 {
					commanderLit = -commanderLit
				}
				// lit -> commander
				result = append(result, cnfAtLeast1(c, []int{-lit, commanderLit})...)
			}
		}
	}

	return result
}

func getBinLength(m int) int {
	len := 1
	for {
		if m <= len {
			break
		}
		m <<= 1
	}
	return len
}
