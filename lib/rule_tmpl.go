package diffence

import (
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

// Run runs rules against input strings
func (r *rule) Run(in string) bool {
	switch r.Type {
	case RuleTypeRegex:
		reg := regexp.MustCompile(r.Pattern)
		return reg.MatchString(in)
	case RuleTypeMatch:
		return strings.Contains(in, r.Pattern)
	}
	return false
}
