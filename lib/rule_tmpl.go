package diffence

import (
	"fmt"
	"regexp"
	"strings"
)

//go:generate gojson -name rule -input ./../test/fixtures/rules/rule.json -o rule.go -pkg diffence -subStruct -tags 'rule'

////////////////////////////////////////////////////////
// https://github.com/michenriksen/gitrob#signature-keys
////////////////////////////////////////////////////////
const (
	// RuleTypeRegex is the regex type for pattern matching
	RuleTypeRegex = "regex"
	// regex: Regular expression matching of part and pattern
	// https://golang.org/pkg/regexp/#Regexp.FindAllString
	// re := regexp.MustCompile("a.")
	// fmt.Println(re.FindAllString("paranormal", -1))
	// matched, err := regexp.MatchString("foo.*", "seafood")
	// fmt.Println(matched, err)

	// RuleTypeMatch is the string match type for pattern matching
	RuleTypeMatch = "match"
	// match: Simple match of part and pattern
	// strings.Contains("seafood", "foo")
	// https://golang.org/pkg/regexp/#Regexp.MatchString
	// https://golang.org/pkg/regexp/#Regexp.Match
)

const (
	// RulePartPath checks the path of the file
	RulePartPath = "path"
	// complete file path
	// Only the file extension

	// RulePartFilename checks the name of the file
	RulePartFilename = "filename"
	// Only the filename
	// path.Base()

	// RulePartExtension checks the extension of the file
	RulePartExtension = "extension"
	// path.Ext()
)

// RuleResult returns the result of the pattern matching
type RuleResult struct {
	Matched bool
	Err     error
}

//////////////
// pattern: (match = string)
// regular expression to match with
//////////////
func (r *rule) Run(in string) RuleResult {
	switch r.Type {
	case RuleTypeRegex:
		return r.runRegex(in)
	case RuleTypeMatch:
		return r.runMatch(in)
	}
	return RuleResult{
		Matched: false,
		Err:     fmt.Errorf("Unrecognised rule type: %s", r.Type),
	}
}

func (r *rule) runRegex(in string) RuleResult {
	var matched bool
	reg, err := regexp.Compile(r.Pattern)
	if err != nil {
		matched = reg.MatchString(in)
	}
	return RuleResult{
		Matched: matched,
		Err:     err,
	}
}

func (r *rule) runMatch(in string) RuleResult {
	return RuleResult{
		Matched: strings.Contains(in, r.Pattern),
		Err:     nil,
	}
}
