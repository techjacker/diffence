package diffence

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
)

// LoadRulesJSONFromPwd reads a rules JSON from a path relative to the process's pwd
func LoadRulesJSONFromPwd(rulesPath string) *[]Rule {
	_, cmd, _, _ := runtime.Caller(0)
	rules, err := LoadRulesJSON(path.Join(path.Dir(cmd), rulesPath))
	if err != nil {
		panic(fmt.Sprintf("Cannot read rule file: %s\n", err))
	}
	return rules
}

// LoadRulesJSON reads a file of JSON rules from the local filesystem
func LoadRulesJSON(filePath string) (*[]Rule, error) {
	rules := &[]Rule{}

	f, err := os.Open(filePath)
	if err != nil {
		return rules, err
	}

	jsonParser := json.NewDecoder(f)
	err = jsonParser.Decode(rules)
	return rules, err
}
