package sudokusolver

func cnfAtLeast1(c CNFInterface, lits []int) {
	c.addClause(lits)
}

func cnfAtMost1(c CNFInterface, lits []int) {
	_cnfAtMost1(c, lits, false)
}

func _cnfAtMost1(c CNFInterface, lits []int, pairwise bool) {
	// cnfAtMost1Pairwise(c, lits)
	if pairwise || len(lits) <= 5 {
		cnfAtMost1Pairwise(c, lits)
		return
	}

	if len(lits) <= 10 {
		cnfAtMost1Commander(c, lits)
		return
	}

	cnfAtMost1Bimander(c, lits)
}

func cnfExactly1(c CNFInterface, lits []int) {
	if len(lits) == 1 {
		c.addLit(lits[0])
		return
	}

	cnfAtMost1(c, lits)
	cnfAtLeast1(c, lits)
}

func cnfAtMost1Pairwise(c CNFInterface, lits []int) {
	n := len(lits)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// each pair can't be true at the same time
			cnfAtLeast1(c, []int{-lits[i], -lits[j]})
		}
	}
}

const COMMANDER_FACTOR = 3

// Will Klieber and Gihwon Kwon. 2007.
// Efficient CNF Encoding for Selecting 1 from N Objects.
func cnfAtMost1Commander(c CNFInterface, lits []int) {
	n := len(lits)
	if n <= 3 {
		_cnfAtMost1(c, lits, true)
		return
	}
	m := (n + COMMANDER_FACTOR - 1) / COMMANDER_FACTOR
	groupLen := (n + m - 1) / m

	groups := make([][]int, m)
	for i := 0; i < m; i++ {
		groups[i] = lits[groupLen*i : min(groupLen*(i+1), n)]
		// 1. At most one variable in a group can be true
		_cnfAtMost1(c, groups[i], true)
	}

	commanders := c.requestLiterals(uint32(m))

	//  2. If the commander variable of a group is false,
	// then none of the variables in the group can be true
	for i, commander := range commanders {
		for _, lit := range groups[i] {
			// -commander -> -lit
			cnfAtLeast1(c, []int{commander, -lit})
		}
	}

	// 3. Exactly one of the commander variables is true
	cnfAtLeast1(c, commanders)
	cnfAtMost1Commander(c, commanders)
}

const BIMANDER_FACTOR = 2

// Nguyen, Van-Hau, and Son T. Mai. 2015.
// A new method to encode the at-most-one constraint into SAT.
func cnfAtMost1Bimander(c CNFInterface, lits []int) {
	n := len(lits)
	m := (n + BIMANDER_FACTOR - 1) / BIMANDER_FACTOR
	groupLen := (n + m - 1) / m
	binLength := getBinLength(uint32(m))

	groups := make([][]int, m)
	for i := 0; i < m; i++ {
		groups[i] = lits[groupLen*i : min(groupLen*(i+1), n)]
		_cnfAtMost1(c, groups[i], true)
	}

	auxVars := c.requestLiterals(binLength)

	for i := 0; i < m; i++ {
		for _, lit := range groups[i] {
			for j, aux := range auxVars {
				commanderLit := aux
				if (i & (1 << j)) == 0 {
					commanderLit = -commanderLit
				}
				// lit -> commander
				cnfAtLeast1(c, []int{-lit, commanderLit})
			}
		}
	}
}

func getBinLength(m uint32) uint32 {
	len := uint32(1)
	for m > len {
		len <<= 1
	}
	return len
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func makeRange(min, max uint32) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = int(min) + i
	}
	return a
}
