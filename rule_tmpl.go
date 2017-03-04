package diffence

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

//go:generate gojson -name Rule -input test/fixtures/rules/rule.json -o rule.go -pkg diffence -subStruct -tags 'Rule'

// https://github.com/michenriksen/gitrob#signature-keys
const (
	// RuleTypeRegex is the regex type for pattern matching
	RuleTypeRegex = "regex"

	// RuleTypeMatch is the string match type for pattern matching
	RuleTypeMatch = "match"
)

const (
	// RulePartPath checks the whole path of the file
	RulePartPath = "path"

	// RulePartFilename checks the name of the file
	RulePartFilename = "filename"

	// RulePartExtension checks the extension of the file
	RulePartExtension = "extension"
)

// Match runs rules against input strings
func (r *Rule) Match(in string) bool {
	in = r.extractPart(in)
	switch r.Type {
	case RuleTypeRegex:
		// make match case insensitive
		reg := regexp.MustCompile("(?i)" + r.Pattern)
		// fmt.Printf("%#v\n", reg)
		// reg.Op = OpBeginLine | OpEndLine
		// fmt.Printf("%s\n", r.Pattern)
		return reg.MatchString(in)
	case RuleTypeMatch:
		// make match case insensitive
		return strings.Contains(strings.ToLower(in), strings.ToLower(r.Pattern))
		// return strings.Contains(in, r.Pattern)
	}
	return false
}

func (r *Rule) extractPart(in string) string {
	switch r.Part {
	case RulePartFilename:
		return path.Base(in)
	case RulePartExtension:
		return path.Ext(in)
	}
	return in
}

// String returns a string representation of the rule
func (r *Rule) String() string {
	return fmt.Sprintf("Caption: %s\n", r.Caption) +
		fmt.Sprintf("Description: %#v\n", r.Description) +
		fmt.Sprintf("Part: %s\n", r.Part) +
		fmt.Sprintf("Pattern: %s\n", r.Pattern) +
		fmt.Sprintf("Type: %s\n\n", r.Type)
}
