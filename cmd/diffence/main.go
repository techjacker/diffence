package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/techjacker/diffence"
)

func main() {

	rPath := flag.String("rules", "", "path to rules in JSON format")
	flag.Parse()

	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		log.Fatalln("The command is intended to work with pipes.")
		return
	}

	var (
		err   error
		rules *[]diffence.Rule
	)

	if len(*rPath) > 0 {
		rules, err = diffence.LoadRulesJSON(*rPath)
	} else {
		rules, err = diffence.LoadDefaultRules()
	}
	if err != nil {
		log.Fatalf("Cannot load rules\n%s", err)
		return
	}

	diff := diffence.DiffChecker{
		Rules:   rules,
		Ignorer: diffence.NewIgnorerFromFile(".secignore"),
	}
	res, err := diff.Check(bufio.NewReader(os.Stdin))
	if err != nil {
		log.Fatalf("Error reading diff\n%s\n", err)
		return
	}

	// log results to STDOUT/STDERR
	logger := log.New(os.Stdout, "", 0x0)
	if res.Matches() > 0 {
		log.SetOutput(os.Stderr)
		defer os.Exit(1)
	}
	res.Log(logger)
}
