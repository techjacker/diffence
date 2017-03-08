package diffence

import (
	"fmt"
	"strings"
)

// List is an interface for adding items to a list
type List interface {
	Push(string)
	// Validate() error
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

// type DiffHeader func() string {
// 	return "static"
// }
// 	String() string
// }

// func (d Diffs) Validate() error {
// 	if len(items) < 1 {
// 		return items, errors.New("Not valid diff content")
// 	}
// }

func extractHeader(in string) (string, error) {
	newLineIndex := strings.Index(in, "\n")
	if newLineIndex < 0 {
		return "", fmt.Errorf("not valid diff content:\n\n%s", in)
	}
	return in[:newLineIndex], nil
}

func extractFilePath(in string) (string, error) {
	pathBIndex := strings.Index(in, pathSep)
	newLineIndex := strings.Index(in, "\n")
	if pathBIndex >= 0 && newLineIndex > pathBIndex {
		return in[pathBIndex+len(pathSep) : newLineIndex], nil
	}
	return "", fmt.Errorf("Not valid diff content:\n%s", in)
}
