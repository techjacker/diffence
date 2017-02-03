package diffence

import "io"

// Results is hash of results matched for each filepath in a git diff
// [filepath] => Rule{rule1, rule2}
type Results map[string][]Rule

// CheckDiffs is a clean syntax, inefficient way of
// finding diffs that match the supplied rules
func CheckDiffs(r io.Reader, rules *[]Rule) (Results, error) {
	res := Results{}

	diffs, err := SplitDiffs(r)
	if err != nil || len(diffs) < 1 {
		return res, err
	}

	for _, d := range diffs {
		res[d.filePath] = []Rule{}
		for _, r := range *rules {
			if r.Match(d.filePath) {
				res[d.filePath] = append(res[d.filePath], r)
			}
		}
	}

	return res, err
}
