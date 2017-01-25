package diffence

import (
	"reflect"
	"testing"
)

func Test_extractFileName(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{

			name: "ExtractFileName()",
			args: args{
				in: []byte(
					"diff --git a/README.md b/README.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				),
			},
			want: []byte("README.md"),
		},
		{

			name: "ExtractFileName()",
			args: args{
				in: []byte(
					"diff --git a/TODO.md b/TODO.md" +
						"\n" +
						"index 82366e3..5fc99b9 100644",
				),
			},
			want: []byte("TODO.md"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractFileName(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractFileName()\nGOT:%s\nWANT:%s", got, tt.want)
			}
		})
	}
}

func Test_extractAddedText(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractAddedText(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractAddedText() = %v, want %v", got, tt.want)
			}
		})
	}
}
