package diffence

import (
	"fmt"
	"strings"
)

const (
	pathSep = " b/"
)

// List is an interface for adding items to a list
type List interface {
	Push(string)
}

// DiffItem is a diff struct for an inidividual file
type DiffItem struct {
	raw   string
	fPath string
	// w      io.Writer
	// fPath fmt.Stringer
	// addedText string
	// match        bool
	// matchedRules []rule
}

// Diff is a list of split diffs
type Diff struct {
	ignorer Matcher
	Items   []DiffItem
	Error   error
}

// Push a diff on to the list
func (d *Diff) Push(s string) {
	fPath, err := extractFilePath(s)
	if err != nil {
		d.Error = err
		return
	}

	// fmt.Printf("%s\n", d.ignorer.Match(fPath))
	// if d.ignorer.Match(fPath) == true {
	// 	return
	// }

	d.Items = append(d.Items, DiffItem{
		raw:   s,
		fPath: fPath,
	})
}

func extractFilePath(in string) (string, error) {
	pathBIndex := strings.Index(in, pathSep)
	newLineIndex := strings.Index(in, "\n")
	if pathBIndex >= 0 && newLineIndex > pathBIndex {
		return in[pathBIndex+len(pathSep) : newLineIndex], nil
	}
	return "", fmt.Errorf("Not valid diff content:\n%s", in)
}
