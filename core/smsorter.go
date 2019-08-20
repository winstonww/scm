package core

import (
	"sort"
)

type lessThan func(sm1, sm2 *SwitchMigration) bool

// multiSorter implements the Sort interface, sorting the changes within.
type SmSorter struct {
	migrations []SwitchMigration
	less       []lessThan
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (s *SmSorter) Sort(migrations []SwitchMigration) {
	s.migrations = migrations
	sort.Sort(s)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessThan) *SmSorter {
	return &SmSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (s *SmSorter) Len() int {
	return len(s.migrations)
}

// Swap is part of sort.Interface.
func (s *SmSorter) Swap(i, j int) {
	s.migrations[i], s.migrations[j] = s.migrations[j], s.migrations[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that discriminates between
// the two items (one is less than the other). Note that it can call the
// less functions twice per call. We could change the functions to return
// -1, 0, 1 and reduce the number of calls for greater efficiency: an
// exercise for the reader.
func (s *SmSorter) Less(i, j int) bool {
	p, q := &s.migrations[i], &s.migrations[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(s.less)-1; k++ {
		less := s.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return s.less[k](p, q)
}
