package diffence

type rule struct {
	Caption     string      `'rule':"caption"`
	Description interface{} `'rule':"description"`
	Part        string      `'rule':"part"`
	Pattern     string      `'rule':"pattern"`
	Type        string      `'rule':"type"`
}
