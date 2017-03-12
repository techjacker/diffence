package diffence

import "io"

// Checker checks diffs for rule violations
type Checker interface {
	Check(io.Reader) (Result, error)
}

// DiffChecker checks an io.Reader for matches against the supplied ruleset
type DiffChecker struct {
	Rules   *[]Rule
	Ignorer Matcher
}

// Check is a clean syntax but memory inefficient
// method for finding diffs that match the supplied rules
// (use an array instead of a map for better performance)
func (dc DiffChecker) Check(r io.Reader) (Result, error) {
	res := Result{
		Matched:      false,
		MatchedRules: make(map[string][]Rule),
	}

	diff := Diff{ignorer: dc.Ignorer}
	err := SplitDiffs(r, &diff)

	for _, d := range diff.Items {
		for _, r := range *dc.Rules {
			if r.Match(d.fPath) {
				res.Matched = true
				if _, ok := res.MatchedRules[d.fPath]; !ok {
					res.MatchedRules[d.fPath] = []Rule{}
				}
				res.MatchedRules[d.fPath] = append(res.MatchedRules[d.fPath], r)
			}
		}
	}

	return res, err
}
