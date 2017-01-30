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
type RuleType string

const (
	// regex: Regular expression matching of part and pattern
	// https://golang.org/pkg/regexp/#Regexp.FindAllString
	// re := regexp.MustCompile("a.")
	// fmt.Println(re.FindAllString("paranormal", -1))
	// matched, err := regexp.MatchString("foo.*", "seafood")
	// fmt.Println(matched, err)
	REGEX RuleType = "regex"
	// match: Simple match of part and pattern
	// strings.Contains("seafood", "foo")
	// https://golang.org/pkg/regexp/#Regexp.MatchString
	// https://golang.org/pkg/regexp/#Regexp.Match
	MATCH RuleType = "match"
)

type RulePart string

const (
	// complete file path
	PATH RulePart = "part"
	// Only the filename
	// path.Base()
	FILENAME = "filename"
	// Only the file extension
	// path.Ext()
	EXTENSION = "extension"
)

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
	case REGEX:
		return r.runRegex(in)
	case MATCH:
		return r.runMatch(in)
	}
	return RuleResult{
		Matched: false,
		Err:     fmt.Errorf("Unrecognised rule type: %s", r.Type),
	}
}

func (r *rule) runRegex(in string) RuleResult {
	reg := regexp.MustCompile(r.Pattern)
	matched, err := reg.MatchString(in)

	return RuleResult{
		Matched: matched,
		Err:     err,
	}
}

func (r *rule) runMatch(in string) RuleResult {
	return RuleResult{
		Matched: strings.Contains(in, r.Pattern),
		Err:     err,
	}
}
