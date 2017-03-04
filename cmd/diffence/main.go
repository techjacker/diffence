package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/techjacker/diffence"
)

const rulesPath = "test/fixtures/rules/gitrob.json"

func main() {

	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		log.Fatalln("The command is intended to work with pipes.")
	}

	diff := diffence.DiffChecker{
		Rules: diffence.LoadRulesJSONFromPwd(rulesPath),
	}
	res, err := diff.Check(bufio.NewReader(os.Stdin))
	if err != nil {
		log.Fatalf("Error reading diff\n%s\n", err)
		return
	}

	matches := res.Matches()
	if matches > 0 {
		i := 1
		fmt.Printf("Diff contains %d offenses\n\n", matches)
		for filename, rule := range res.MatchedRules {
			fmt.Printf("------------------\n")
			fmt.Printf("Violation %d\n", i)
			fmt.Printf("File: %s\n", filename)
			fmt.Printf("Reason: %#v\n\n", rule[0].Caption)
			i++
		}
		os.Exit(1)
		return
	}
	fmt.Printf("Diff contains NO offenses\n\n")
	os.Exit(0)
}
