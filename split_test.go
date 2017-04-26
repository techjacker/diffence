package diffence

import (
	"io"
	"testing"
)

func TestSplitDiffs(t *testing.T) {
	type args struct {
		r io.Reader
	}

	tests := []struct {
		name string
		args args
		want []DiffItem
	}{
		// {
		// 	name: "SplitDiffs()",
		// 	args: args{r: getFixtureFile("test/fixtures/diffs/single_commit_id.diff")},
		// 	want: []DiffItem{
		// 		{
		// 			fPath:  "README.md",
		// 			commit: "17949087b8e0c9179345e8dbb7b6705b49c93c77 Adds results logger",
		// 		},
		// 	},
		// },
		{
			name: "SplitDiffs()",
			args: args{r: getFixtureFile("test/fixtures/diffs/single.diff")},
			want: []DiffItem{
				{fPath: "README.md"},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{r: getFixtureFile("test/fixtures/diffs/multi.diff")},
			want: []DiffItem{
				{fPath: "TODO.md"},
				{fPath: "systemdlogger/aws.py"},
				{fPath: "systemdlogger/cloudwatch.py"},
				{fPath: "tests/fixtures/config.json"},
				{fPath: "tests/test_aws.py"},
				{fPath: "tests/test_cloudwatch.py"},
				{fPath: "tests/test_runner_integration.py"},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{r: getFixtureFile("test/fixtures/diffs/logp.truncated.diff")},
			want: []DiffItem{
				{fPath: "README.md"},
				{fPath: "TODO.md"},
				{fPath: "check.go"},
				{fPath: "results.go"},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{r: getFixtureFile("test/fixtures/diffs/longline.diff")},
			want: []DiffItem{
				{fPath: "web/src/main/resources/static/js/menu.js"},
				{fPath: "web/src/main/resources/static/js/vendor/jquery-1.11.1.min.js"},
				{fPath: "web/src/main/resources/templates/layout.vm"},
			},
		},
		{
			name: "SplitDiffs()",
			args: args{r: getFixtureFile("test/fixtures/diffs/logp.diff")},
			want: generateWantDiffFromFiles("test/fixtures/diffs/expected/logp.diff.filepaths.txt"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Diff{}
			err := SplitDiffs(tt.args.r, &d)
			if err != nil {
				t.Logf("SplitDiffs =%d", len(d.Items))
				t.Fatalf("SplitDiffs threw error %#v", err)
			}
			for i, di := range d.Items {
				// fmt.Printf("%v\n", di)
				equals(t, tt.want[i].fPath, di.fPath)
				equals(t, tt.want[i].commit, di.commit)
			}
		})
	}
}
