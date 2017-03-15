package diffence

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestCheckDiffs(t *testing.T) {
	type args struct {
		r       io.Reader
		rules   *[]Rule
		ignorer Matcher
	}
	ruleSingle := getRuleFile("test/fixtures/rules/rules.json")
	ruleMulti := getRuleFile("test/fixtures/rules/rules_multi.json")
	rulesExtended := getRuleFile("test/fixtures/rules/rules_extended_regex.json")

	tests := []struct {
		name string
		args args
		want Result
	}{
		{
			name: "Recognises an offensive diff - sqlpassword.diff",
			args: args{
				r:     getFixtureFile("test/fixtures/diffs/sqlpassword.diff"),
				rules: rulesExtended,
			},
			want: Result{
				Matched: true,
				MatchedRules: MatchedRules{
					"web/src/main/resources/db/migration/V0_4__AdminPassword.sql": *rulesExtended,
					"web/src/main/resources/db/migration/V0_2__SeedData.sql":      []Rule{(*rulesExtended)[0]},
				},
			},
		},
		{
			name: "Recognises an offensive diff",
			args: args{
				r:     getFixtureFile("test/fixtures/diffs/single_fail.diff"),
				rules: ruleSingle,
			},
			want: Result{
				Matched: true,
				MatchedRules: MatchedRules{
					"path/to/password.txt": *ruleSingle,
				},
			},
		},
		{
			name: "Recognises an offensive diff - single_fail.diff",
			args: args{
				r:     getFixtureFile("test/fixtures/diffs/single_fail.diff"),
				rules: ruleMulti,
			},
			want: Result{
				Matched: true,
				MatchedRules: MatchedRules{
					"path/to/password.txt": *ruleSingle,
				},
			},
		},
		{
			name: "Recognises an offensive diff - multi_fail.diff",
			args: args{
				r:     getFixtureFile("test/fixtures/diffs/multi_fail.diff"),
				rules: ruleMulti,
			},
			want: Result{
				Matched: true,
				MatchedRules: MatchedRules{
					"path/to/password.txt": []Rule{(*ruleMulti)[0]},
					"another/file/aws.pem": []Rule{(*ruleMulti)[1]},
				},
			},
		},
		{
			name: "Recognises an offensive diff - but excludes file in ignorer",
			args: args{
				r:       getFixtureFile("test/fixtures/diffs/multi_fail.diff"),
				rules:   ruleMulti,
				ignorer: Ignorer{patterns: []string{"another/file/*.pem"}},
			},
			want: Result{
				Matched: true,
				MatchedRules: MatchedRules{
					"path/to/password.txt": []Rule{(*ruleMulti)[0]},
				},
			},
		},
		{
			name: "Recognises non diff text",
			args: args{
				r:     bytes.NewReader([]byte("not a diff")),
				rules: ruleMulti,
			},
			want: Result{
				Matched:      false,
				MatchedRules: MatchedRules{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := DiffChecker{Rules: tt.args.rules, Ignorer: tt.args.ignorer}
			got, _ := dc.Check(tt.args.r)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckDiffs(): %s\n\n got:%#v\n\nwant: %#v", tt.name, got, tt.want)
			}
		})
	}
}
