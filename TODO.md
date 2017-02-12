add results tests

add logger
	- Logger interface fn arg - diff backends possible
		- see in docs - default logger to be used
		- alternative = return string (for email/alt export)
		- cloudwatch/loggly alts  etc
	- CLI logger
		- add color ([OK] [ERR])
		- fmt.Fprint(os.Stdout, "Hello ", 23, "\n")

Speed up - pauses on stin pipe
	// r := bytes.NewReader([]byte("hello world"))
 	writer := bufio.NewWriter(outfile)
    defer writer.Flush()

Add context.cancel
	- sigterm/cancel -> stop CLI


-----------------------------------------------------------
write git hook integrations:
	- yelp's pre-commit
	- overcommit

-----------------------------------------------------------
pull request gojson lib - add description flag
	- re-enable lint githook (disable for that file)

-----------------------------------------------------------
github integration - HTTP server
- set up fission on kubernetes

-----------------------------------------------------------
-----------------------------------------------------------
config
	- rule file locations
		- fs
		- http
	- rule file for:
		1. filepaths
		2. added lines

-----------------------------------------------------------
-----------------------------------------------------------
-----------------------------------------------------------

Add -config flag to CLI
	- override default = gitrob rules

	-----------------------------------------------------------
	## JSON Configuration

	By default the CLI tool looks for config.json in $PWD. You can specify a custom location with the `config` flag, eg:

	```Shell
	diffence -config ./test/fixtures/config.json
	```


	#### Example JSON Config
	```json
	{
		"rules": {
			"jsonPath": "./test/fixtures/rules.json"
		}
	}
	```
add option to add multiple rules files
	- bufio.MultiReader

-----------------------------------------------------------
-----------------------------------------------------------
Add benchmarking

-----------------------------------------------------------
Add concurrency
	- each NewDiffItem = new go routine
	- buffer to max 100?
		- run benchmark tests (add to tests + makefile)

-----------------------------------------------------------
Perf
	- replace map with arrays (not slices) for diffs + rules
	- move SplitDiffs() into check()




-----------------------------------------------------------


-----------------------------------------------------------
add build task -> convert JSON rules into golang struct
	- make part of dockerfile?

-----------------------------------------------------------


-----------------------------------------------------------
set up realize - live reload run tests etc
