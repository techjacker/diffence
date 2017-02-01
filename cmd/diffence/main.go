package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	df "github.com/techjacker/diffence"
)

func main() {
	// r := bytes.NewReader([]byte("hello world"))
	gitrobRuleFile := "../../test/fixtures/rules/gitrob.json"
	_, cmd, _, _ := runtime.Caller(0)
	rules, err := df.ReadRulesFromFile(path.Join(path.Dir(cmd), gitrobRuleFile))
	if err != nil {
		log.Fatalf("\nCannot read rule file: %s\n", err)
		return
	}

	res, err := df.CheckDiffs(os.Stdin, rules)
	if err != nil {
		log.Fatalf("\nError reading diff\n%s\n", err)
		return
	}

	dirty := false
	for k, v := range res {
		if len(v) > 0 {
			dirty = true
			fmt.Printf("\nFile %s violates %d rules:\n", k, len(v))
			for _, r := range v {
				fmt.Printf("\n%s\n", r.String())
			}
		}
	}

	if dirty == false {
		fmt.Printf("\nDiff contains no offenses\n\n")
		os.Exit(0)
	}
	// dirty == true
	os.Exit(1)
}
