package diffence

import (
	"bufio"
	"io"
)

type diffItem struct {
	data         []byte
	filename     []byte
	addedText    []byte
	match        bool
	matchedRules []rule
}

// Differ creates diffItems from a raw git diff text input
type Differ interface {
	// split(string) []string
	Parse(io.Reader) []diffItem
}

// NewDiffer is a Differ factory
func NewDiffer() Differ {
	return &diff{}
}

type diff struct{}

func (d diff) Parse(r io.Reader) []diffItem {

	items := []diffItem{}

	// Default scanner is bufio.ScanLines. Lets use ScanWords.
	// Could also use a custom function of SplitFunc type
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Bytes()
		items = append(items, diffItem{word})
	}

	return items
}
