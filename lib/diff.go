package diffence

type diffItem struct {
	rawText      string
	filename     string
	addedText    string
	match        bool
	matchedRules []rule
}

// Differ creates diffItems from a raw git diff text input
type Differ interface {
	split(string) []string
	parse([]string) []diffItem
	Run() error
}

// NewDiffer is a Differ factory
func NewDiffer(in string) Differ {
	return &diff{rawText: in}
}

type diff struct {
	rawText string
}

func (d diff) split(s string) []string {
	return []string{}
}

func (d diff) parse(sArr []string) []diffItem {
	return []diffItem{}
}

func (d diff) Run() error {
	return nil
}
