package diffence

import "io"

// MatchedRules is slice of matched rules for each file in diff
// [filepath] => Rule{rule1, rule2}
type MatchedRules map[string][]Rule

// Results is a slice of Result structs
type Results []Result

// Result compiles the results of matched rules for a diff
type Result struct {
	// Have any of the files matches against the rules?
	Matched      bool
	MatchedRules MatchedRules
}

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

	diffs, err := SplitDiffs(r)
	if err != nil || len(diffs) < 1 {
		return res, err
	}

	for _, d := range diffs {
		for _, r := range *dc.Rules {
			if r.Match(d.filePath) {
				res.Matched = true
				if _, ok := res.MatchedRules[d.filePath]; !ok {
					res.MatchedRules[d.filePath] = []Rule{}
				}
				res.MatchedRules[d.filePath] = append(res.MatchedRules[d.filePath], r)
			}
		}
	}

	return res, err
}
