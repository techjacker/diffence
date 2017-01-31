-----------------------------------------------------------
Add test coverage to Makefile

-----------------------------------------------------------
Add benchmarking

-----------------------------------------------------------
Add concurrency
	- each NewDiffItem = new go routine
	- buffer to max 100?
		- run benchmark tests (add to tests + makefile)

-----------------------------------------------------------
Perf
	- replace map with arrays for diffs + rules
	- move SplitDiffs() into check()




-----------------------------------------------------------
-----------------------------------------------------------
github integration - HTTP server
- set up fission on kubernetes

-----------------------------------------------------------
-----------------------------------------------------------

return just error from main function (vs bool, error)

-----------------------------------------------------------

add logger

-----------------------------------------------------------
add build task -> convert JSON rules into golang struct
	- make part of dockerfile?

-----------------------------------------------------------
add multiple rules files to be inputed
	bufio.MultiReader

-----------------------------------------------------------
set up realize - live reload run tests etc
