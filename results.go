package diffence

// MatchedRules is slice of matched rules for each file in diff
// [fPath] => Rule{rule1, rule2}
type MatchedRules map[string][]Rule

// Results is a slice of Result structs
type Results []Result

// Matches returns the number of diffs which had at least one file match against any rules
func (r Results) Matches() int {
	x := 0
	for _, v := range r {
		if v.Matches() > 0 {
			x++
		}
	}
	return x
}

// Result compiles the results of matched rules for a diff
type Result struct {
	// Have any of the files matches against the rules?
	Matched      bool
	MatchedRules MatchedRules
}

// Matches returns the number of files in the diff that matched against any of the rules
func (r Result) Matches() int {
	return len(r.MatchedRules)
}

// Log prints the results of ma
func (r Result) Log(l Logger) {
	matches := r.Matches()
	if matches < 1 {
		l.Printf("Diff contains NO offenses\n\n")
		return
	}
	i := 1
	l.Printf("Diff contains %d offenses\n\n", matches)
	for filename, rule := range r.MatchedRules {
		l.Printf("------------------\n")
		l.Printf("Violation %d\n", i)
		l.Printf("File: %s\n", filename)
		l.Printf("Reason: %#v\n\n", rule[0].Caption)
		i++
	}
}
