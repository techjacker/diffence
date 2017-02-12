package diffence

// Rule is the format for supplying rules to Diffence to check diffs against
type Rule struct {
	Caption     string      `'Rule':"caption"`
	Description interface{} `'Rule':"description"`
	Part        string      `'Rule':"part"`
	Pattern     string      `'Rule':"pattern"`
	Type        string      `'Rule':"type"`
}
