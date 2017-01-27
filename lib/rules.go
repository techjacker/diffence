package diffence

import (
	"encoding/json"
	"os"
)

func readRulesFromFile(filePath string) (*[]rule, error) {
	rules := &[]rule{}

	f, err := os.Open(filePath)
	if err != nil {
		return rules, err
	}

	jsonParser := json.NewDecoder(f)
	err = jsonParser.Decode(rules)
	return rules, err
}
