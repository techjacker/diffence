package diffence

import (
	"fmt"
	"strings"
)

// List is an interface for adding items to a list
type List interface {
	Push(string)
}

// DiffItem is a diff struct for an inidividual file
type DiffItem struct {
	raw      string
	filePath string
	// w      io.Writer
	// filePath fmt.Stringer
	// addedText string
	// match        bool
	// matchedRules []rule
}

// Diff is a list of split diffs
type Diff struct {
	Items []DiffItem
	Error error
}

// Push a diff on to the list
func (d *Diff) Push(s string) {
	filePath, err := extractFilePath(s)
	if err != nil {
		d.Error = err
		return
	}

	// if shouldIgnore(filePath) == true {
	// 	continue
	// }

	d.Items = append(d.Items, DiffItem{
		raw:      s,
		filePath: filePath,
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
