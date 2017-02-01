package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	df "github.com/techjacker/diffence"
)

const rulesPath = "../../test/fixtures/rules/gitrob.json"

func main() {
	_, cmd, _, _ := runtime.Caller(0)
	rules, err := df.ReadRulesFromFile(path.Join(path.Dir(cmd), rulesPath))
	if err != nil {
		log.Fatalf("Cannot read rule file: %s\n", err)
		return
	}

	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		log.Fatalln("The command is intended to work with pipes.")
	}

	res, err := df.CheckDiffs(bufio.NewReader(os.Stdin), rules)
	if err != nil {
		log.Fatalf("Error reading diff\n%s\n", err)
		return
	}

	dirty := false
	for k, v := range res {
		if len(v) > 0 {
			dirty = true
			fmt.Printf("File %s violates %d rules:\n", k, len(v))
			for _, r := range v {
				fmt.Printf("\n%s\n", r.String())
			}
		}
	}

	if dirty == false {
		fmt.Printf("Diff contains no offenses\n\n")
		os.Exit(0)
	}
	// dirty == true
	os.Exit(1)
}
