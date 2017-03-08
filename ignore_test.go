package diffence

import (
	"reflect"
	"testing"
)

func Test_getFileContents(t *testing.T) {
	type args struct {
		fPath string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFileContents(tt.args.fPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFileContents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIgnorer_Match(t *testing.T) {
	type fields struct {
		patterns []string
	}
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Ignorer{
				patterns: tt.fields.patterns,
			}
			if got := i.Match(tt.args.in); got != tt.want {
				t.Errorf("Ignorer.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
