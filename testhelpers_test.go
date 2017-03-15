package diffence

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func generateWantDiffFromFiles(fPathPath string) []DiffItem {
	want := []DiffItem{}
	fPathFile, _ := os.Open(fPathPath)
	r := io.MultiReader(
		fPathFile,
	)
	buffer := bytes.NewBuffer(make([]byte, 0))
	i := 0
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		buffer.Write(scanner.Bytes())
		fPath := buffer.String()
		want = append(want, DiffItem{fPath: fPath})
		i++
		buffer.Reset()
	}
	return want
}

func getFixtureFile(filename string) io.Reader {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func getRuleFile(filename string) *[]Rule {
	rules, err := LoadRulesJSON(filename)
	if err != nil {
		panic(err)
	}
	return rules
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
