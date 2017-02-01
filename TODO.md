stop CLI hanging on empty input from stdin

-----------------------------------------------------------
github integration - HTTP server
- set up fission on kubernetes

-----------------------------------------------------------
-----------------------------------------------------------
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

export rule
	-> rename to Rule
	- re-enable lint githook (disable for that file)
	- pull request gojson lib - add description flag

-----------------------------------------------------------


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
add logger

-----------------------------------------------------------
add build task -> convert JSON rules into golang struct
	- make part of dockerfile?

-----------------------------------------------------------


-----------------------------------------------------------
set up realize - live reload run tests etc
