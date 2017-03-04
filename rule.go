package diffence

// Rule defines a pattern to match against a diff
type Rule struct {
	Caption     string      `'Rule':"caption"`
	Description interface{} `'Rule':"description"`
	Part        string      `'Rule':"part"`
	Pattern     string      `'Rule':"pattern"`
	Type        string      `'Rule':"type"`
}
