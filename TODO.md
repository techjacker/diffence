Convert extract to lexer
	- make accept io.Reader (not bytes as arg)

Add more tests
	- extract_test (/lexer_test)
		- filename
		- addedText
	- diff_test
		- multi diff test expectations
		- add more diff fixture files

-----------------------------------------------------------
Make rules struct private
	- add getter for channels to access

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
