package diffence

import "io"

// Checker checks diffs for rule violations
type Checker interface {
	Check(io.Reader) (Result, error)
}

// DiffChecker checks an io.Reader for matches against the supplied ruleset
type DiffChecker struct {
	Rules *[]Rule
}

// Check is a clean syntax but memory inefficient
// method for finding diffs that match the supplied rules
// (use an array instead of a map for better performance)
func (dc DiffChecker) Check(r io.Reader) (Result, error) {
	res := Result{
		Matched:      false,
		MatchedRules: make(map[string][]Rule),
	}
	diff := Diff{}
	err := SplitDiffs(r, &diff)

	for _, d := range diff.Items {
		for _, r := range *dc.Rules {
			if r.Match(d.filepath) {
				res.Matched = true
				if _, ok := res.MatchedRules[d.filepath]; !ok {
					res.MatchedRules[d.filepath] = []Rule{}
				}
				res.MatchedRules[d.filepath] = append(res.MatchedRules[d.filepath], r)
			}
		}
	}

	return res, err
}
