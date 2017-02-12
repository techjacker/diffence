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

	diff := diffence.DiffChecker{diffence.LoadRulesJSONFromPwd(rulesPath)}
	res, err := diff.Check(bufio.NewReader(os.Stdin))
	if err != nil {
		log.Fatalf("Error reading diff\n%s\n", err)
		return
	}

	if res.Matched == true {
		fmt.Printf("Diff contains offenses\n\n")
		os.Exit(1)
	}
	fmt.Printf("Diff contains NO offenses\n\n")
	os.Exit(0)
}
