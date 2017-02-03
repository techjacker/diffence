package diffence

import (
	"encoding/json"
	"os"
)

// ReadRulesFromFile reads rules from the local filesystem
func ReadRulesFromFile(filePath string) (*[]Rule, error) {
	rules := &[]Rule{}

	f, err := os.Open(filePath)
	if err != nil {
		return rules, err
	}

	jsonParser := json.NewDecoder(f)
	err = jsonParser.Decode(rules)
	return rules, err
}
