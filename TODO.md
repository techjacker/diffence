

d.getFilename() -> getPath()
	- move out of struct into separate function



-----------------------------------------------------------
Add concurrency
	- each NewDiffItem = new go routine
	- buffer to max 100?
		- run benchmark tests (add to tests + makefile)

-----------------------------------------------------------

return just error from main function (vs bool, error)

add logger

-----------------------------------------------------------
add build task -> convert JSON rules into golang struct
	- make part of dockerfile?

add multiple rules files to be inputed
	bufio.MultiReader

set up realize - live reload run tests etc
