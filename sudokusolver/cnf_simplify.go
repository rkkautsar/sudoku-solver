package sudokusolver

import (
	"log"
)

type SimplifyOptions struct {
	disablePureLiteralElimination bool
}

func (c *CNF) Simplify(options SimplifyOptions) {
	// c.Print(os.Stderr)
	c.watchers = make([][]int, 2*c.nbVar)
	posCount := make([]int, c.nbVar)
	negCount := make([]int, c.nbVar)
	// grow litLookup
	c.litLookup = append(c.litLookup, make([]uint8, int(c.nbVar)-len(c.litLookup))...)

	pure := 0
	for i, clause := range c.Clauses {
		for j := 0; j < min(2, len(clause)); j++ {
			idx := c.litToIdx(clause[j])
			c.watchers[idx] = append(c.watchers[idx], i)
		}

		if options.disablePureLiteralElimination {
			continue
		}

		for j := 0; j < len(clause); j++ {
			if clause[j] > 0 {
				posCount[clause[j]-1]++
			} else {
				negCount[-clause[j]-1]++
			}
		}
	}

	queue := make([]int, 0, len(c.lits))

	if !options.disablePureLiteralElimination {
		for i := 1; i <= int(c.nbVar); i++ {
			if c.lookupTrue(i) || c.lookupTrue(-i) {
				continue
			}
			pos := posCount[i-1]
			neg := negCount[i-1]
			if pos == 0 && neg > 0 {
				// log.Println("pure", -i)
				c.addLit(-i)
				pure++
			} else if pos > 0 && neg == 0 {
				// log.Println("pure", i)
				c.addLit(i)
				pure++
			}
		}
	}

	// log.Println("Eliminating", pure, "pure literals")

	// log.Println("adding", c.lits)
	queue = append(queue, c.lits...)

	c.propagateAll(queue)
	// lenBefore := len(c.Clauses)
	c.filterSatisfiedClause()
	// log.Println("Filtered", lenBefore-len(c.Clauses), "clauses")
}

func (c *CNFParallel) Simplify(options SimplifyOptions) {
	c.CNF.Simplify(options)
}

func (c *CNF) propagateAll(queue []int) {
	for len(queue) > 0 {
		lit := queue[0]
		queue = queue[1:]
		c.propagate(queue, lit)
	}
}

func (c *CNF) propagate(queue []int, lit int) {
	// log.Println("propagating", lit, "queue=", len(queue))

	// invariant: "unit" clause should be satisfied by its 0th index lit

	// remove -lit from clauses
	negLit := c.litToIdx(-lit)
	// log.Println(c.watchers[negLit])
	for _, clauseIdx := range c.watchers[negLit] {
		// c.removeLit(clauseIdx, -lit)
		clause := c.Clauses[clauseIdx]

		// put other watch in 0-th index
		if clause[0] == -lit {
			clause[0], clause[1] = clause[1], clause[0]
		}
		if clause[1] != -lit {
			continue
		}

		// log.Println("lookupTrue", clause[0], c.lookupTrue(clause[0]))
		if c.lookupTrue(clause[0]) {
			// already satisfied
			continue
		}

		// log.Println("finding replacement for", clause[1], "in", clause)
		// find replacement
		for j := 2; j < len(clause); j++ {
			// log.Println("considering", clause[j], "as replacement")
			if c.lookupTrue(-clause[j]) {
				continue
			}
			idx := c.litToIdx(clause[j])
			c.watchers[idx] = append(c.watchers[idx], clauseIdx)

			clause[j], clause[1] = clause[1], clause[j]
			break
			// no need to remove from negLit watchers
		}

		if clause[1] != -lit {
			// replaced
			continue
		}

		// has to be satisfied by 0th lit
		if c.litLookup[getLitLookupIdx(clause[0])] != unassigned {
			log.Println(c.lits)
			log.Println(clause, "has to be satisfied by", clause[0], "but it's", c.litLookup[getLitLookupIdx(clause[0])])
			for _, l := range clause {
				log.Println("litLookup", l, c.litLookup[getLitLookupIdx(l)])
			}
			panic("UNSAT")
		}

		// log.Println("adding", clause[0])
		c.addLit(clause[0])
		queue = append(queue, clause[0])
	}
	c.watchers[negLit] = c.watchers[negLit][:0]
}

func (c *CNF) filterSatisfiedClause() {
	filtered := c.Clauses[:0]
	for _, clause := range c.Clauses {
		if !c.isSatisfied(clause) {
			filtered = append(filtered, clause)
		}
	}
	c.Clauses = filtered
}

func (c *CNF) isSatisfied(clause []int) bool {
	for _, l := range clause {
		if c.lookupTrue(l) {
			return true
		}
	}
	return false
}

func (c *CNF) litToIdx(lit int) int {
	if lit < 0 {
		return (-lit-1)<<1 + 1
	}
	return (lit - 1) << 1
}
