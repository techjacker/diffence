package diffence

type Rule struct {
	Caption     string      `'Rule':"caption"`
	Description interface{} `'Rule':"description"`
	Part        string      `'Rule':"part"`
	Pattern     string      `'Rule':"pattern"`
	Type        string      `'Rule':"type"`
}
