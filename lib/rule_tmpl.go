package diffence

//go:generate gojson -name rule -input ./../test/fixtures/rules/rule.json -o rule.go -pkg diffence -subStruct -tags 'rule'

//
// https://github.com/michenriksen/gitrob#signature-keys
//
//////////////
// part
//////////////
// path: The complete file path

// filename: Only the filename
// path.Base()

// extension: Only the file extension
// path.Ext()

//////////////
// type
//////////////
// match: Simple match of part and pattern
// strings.Contains("seafood", "foo")

// regex: Regular expression matching of part and pattern

//////////////
// pattern: The value or regular expression to match with
//////////////

// func(r rule) () {
