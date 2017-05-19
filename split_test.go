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
		{
			name: "SplitDiffs()",
			args: args{r: getFixtureFile("test/fixtures/diffs/single_commit_id.diff")},
			want: []DiffItem{
				{
					fPath:  "README.md",
					commit: "17949087b8e0c9179345e8dbb7b6705b49c93c77",
				},
				{
					fPath:  "check.go",
					commit: "7794f2a7e0c35774f531a74280534374075a9c9e",
				},
				{
					fPath:  "check_test.go",
					commit: "7794f2a7e0c35774f531a74280534374075a9c9e",
				},
			},
		},
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
			want: []DiffItem{
				{
					fPath:  "README.md",
					commit: "6f41b9b0a8150e165cd297ae3e00129766cf8a9b",
				},
				{
					fPath:  "TODO.md",
					commit: "6f41b9b0a8150e165cd297ae3e00129766cf8a9b",
				},
				{
					fPath:  "check.go",
					commit: "6f41b9b0a8150e165cd297ae3e00129766cf8a9b",
				},
				{
					fPath:  "results.go",
					commit: "6f41b9b0a8150e165cd297ae3e00129766cf8a9b",
				},
				{
					fPath:  "LICENSE",
					commit: "50bf1cdde42823e11e78a1026e3a7cfc7bc78e2f",
				},
				{
					fPath:  ".realize/realize.yaml",
					commit: "bf0a0c7499036872255fb6591ad57557d7ec375a",
				},
				{
					fPath:  "TODO.md",
					commit: "bf0a0c7499036872255fb6591ad57557d7ec375a",
				},
				{
					fPath:  "check.go",
					commit: "bf0a0c7499036872255fb6591ad57557d7ec375a",
				},
				{
					fPath:  "check_test.go",
					commit: "bf0a0c7499036872255fb6591ad57557d7ec375a",
				},
			},
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

			equals(t, len(tt.want), len(d.Items))
			for i, w := range tt.want {
				equals(t, w.fPath, d.Items[i].fPath)
				equals(t, w.commit, d.Items[i].commit)
			}
		})
	}
}
